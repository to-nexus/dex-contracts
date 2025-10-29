// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.28;

import {ERC1967Proxy} from "@openzeppelin-contracts-5.2.0/proxy/ERC1967/ERC1967Proxy.sol";
import {IERC20} from "@openzeppelin-contracts-5.2.0/token/ERC20/IERC20.sol";
import {Test} from "forge-std/Test.sol";

import {BuyBot} from "../src/BuyBot.sol";
import {CrossDexImpl} from "../src/CrossDexImpl.sol";
import {CrossDexRouter} from "../src/CrossDexRouter.sol";
import {MarketImpl} from "../src/MarketImpl.sol";
import {PairImpl} from "../src/PairImpl.sol";
import {IPair} from "../src/interfaces/IPair.sol";

contract MockERC20 is Test {
    string public name;
    string public symbol;
    uint8 public decimals;

    mapping(address => uint256) public balanceOf;
    mapping(address => mapping(address => uint256)) public allowance;

    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(address indexed owner, address indexed spender, uint256 value);

    constructor(string memory _name, string memory _symbol, uint8 _decimals) {
        name = _name;
        symbol = _symbol;
        decimals = _decimals;
    }

    function mint(address to, uint256 amount) external {
        balanceOf[to] += amount;
        emit Transfer(address(0), to, amount);
    }

    function transfer(address to, uint256 amount) external returns (bool) {
        balanceOf[msg.sender] -= amount;
        balanceOf[to] += amount;
        emit Transfer(msg.sender, to, amount);
        return true;
    }

    function approve(address spender, uint256 amount) external returns (bool) {
        allowance[msg.sender][spender] = amount;
        emit Approval(msg.sender, spender, amount);
        return true;
    }

    function transferFrom(address from, address to, uint256 amount) external returns (bool) {
        allowance[from][msg.sender] -= amount;
        balanceOf[from] -= amount;
        balanceOf[to] += amount;
        emit Transfer(from, to, amount);
        return true;
    }
}

contract BuyBotTest is Test {
    BuyBot public buyer;
    BuyBot public buyerWithWETH;
    CrossDexImpl public crossDex;
    CrossDexRouter public router;
    MarketImpl public market;
    PairImpl public pair;
    PairImpl public wethPair;

    MockERC20 public quoteToken;
    MockERC20 public baseToken;

    address public owner;
    address public user;
    address public feeCollector;
    address public buyerRole;
    address public recipient;
    address public managerRole;

    uint256 public constant MIN_ORDER_AMOUNT = 100e18;
    uint256 public constant TICK_SIZE = 1e18;
    uint256 public constant LOT_SIZE = 1e18;

    function setUp() public {
        owner = makeAddr("owner");
        user = makeAddr("user");
        feeCollector = makeAddr("feeCollector");
        buyerRole = makeAddr("buyerRole");
        recipient = makeAddr("recipient");
        managerRole = makeAddr("managerRole");

        // Deploy mock tokens
        quoteToken = new MockERC20("USDC", "USDC", 18);
        baseToken = new MockERC20("TOKEN", "TOKEN", 18);

        // Deploy implementations
        CrossDexImpl crossDexImpl = new CrossDexImpl();
        CrossDexRouter routerImpl = new CrossDexRouter();
        MarketImpl marketImpl = new MarketImpl();
        PairImpl pairImpl = new PairImpl();

        // Deploy CrossDex proxy
        bytes memory crossDexInitData = abi.encodeCall(
            CrossDexImpl.initialize, (owner, address(routerImpl), 10, 100, 50, address(marketImpl), address(pairImpl))
        );
        ERC1967Proxy crossDexProxy = new ERC1967Proxy(address(crossDexImpl), crossDexInitData);
        crossDex = CrossDexImpl(address(crossDexProxy));

        router = CrossDexRouter(payable(crossDex.ROUTER()));

        // Create market
        vm.prank(owner);
        address marketAddress = crossDex.createMarket(owner, address(quoteToken), feeCollector, 30);
        market = MarketImpl(marketAddress);

        // Create pair
        vm.prank(owner);
        address pairAddress = market.createPair(address(baseToken), TICK_SIZE, LOT_SIZE);
        pair = PairImpl(pairAddress);

        // Deploy BuyBot for normal token (0 interval = no delay)
        buyer = new BuyBot(owner, address(router), MIN_ORDER_AMOUNT, 0, recipient, buyerRole, managerRole);

        // Setup initial liquidity (sell orders)
        _setupLiquidity();

        // Setup WETH market and pair
        _setupWETHMarket();
    }

    function _setupWETHMarket() internal {
        // Use the same market, just create a WETH pair (WETH as BASE, USDC as QUOTE)
        // Get WETH address from router
        address wethAddress = address(router.CROSS());

        // Create WETH/USDC pair in the same market
        vm.prank(owner);
        address wethPairAddress = market.createPair(wethAddress, TICK_SIZE, LOT_SIZE);
        wethPair = PairImpl(wethPairAddress);

        // Deploy BuyBot for WETH (0 interval = no delay)
        buyerWithWETH = new BuyBot(owner, address(router), MIN_ORDER_AMOUNT, 0, recipient, buyerRole, managerRole);

        // Setup liquidity for WETH pair
        address seller = makeAddr("wethSeller");
        vm.deal(seller, 1000 ether);

        vm.startPrank(seller);
        // Submit sell orders using native ETH
        router.submitSellLimit{value: 100 ether}(
            address(wethPair),
            10e18, // price: 10 USDC
            100 ether, // amount: 100 ETH
            IPair.LimitConstraints.GOOD_TILL_CANCEL,
            [uint256(0), uint256(0)],
            0
        );

        router.submitSellLimit{value: 100 ether}(
            address(wethPair),
            11e18, // price: 11 USDC
            100 ether,
            IPair.LimitConstraints.GOOD_TILL_CANCEL,
            [uint256(10e18), uint256(0)],
            0
        );

        router.submitSellLimit{value: 100 ether}(
            address(wethPair),
            12e18, // price: 12 USDC
            100 ether,
            IPair.LimitConstraints.GOOD_TILL_CANCEL,
            [uint256(11e18), uint256(0)],
            0
        );
        vm.stopPrank();
    }

    function _setupLiquidity() internal {
        address seller = makeAddr("seller");

        // Mint tokens to seller
        baseToken.mint(seller, 1000e18);

        // Create sell orders
        vm.startPrank(seller);
        baseToken.approve(address(router), 1000e18);

        // Submit sell limit orders at different prices
        router.submitSellLimit(
            address(pair),
            10e18, // price: 10 USDC
            100e18, // amount: 100 TOKEN
            IPair.LimitConstraints.GOOD_TILL_CANCEL,
            [uint256(0), uint256(0)],
            0
        );

        router.submitSellLimit(
            address(pair),
            11e18, // price: 11 USDC
            100e18, // amount: 100 TOKEN
            IPair.LimitConstraints.GOOD_TILL_CANCEL,
            [uint256(10e18), uint256(0)],
            0
        );

        router.submitSellLimit(
            address(pair),
            12e18, // price: 12 USDC
            100e18, // amount: 100 TOKEN
            IPair.LimitConstraints.GOOD_TILL_CANCEL,
            [uint256(11e18), uint256(0)],
            0
        );

        vm.stopPrank();
    }

    function test_InitialState() public view {
        assertEq(buyer.owner(), owner);
        assertEq(address(buyer.router()), address(router));
        assertEq(buyer.minOrderAmount(), MIN_ORDER_AMOUNT);
    }

    function test_RevertWhen_BuyerIsZeroAddress() public {
        vm.expectRevert(abi.encodeWithSelector(BuyBot.BuyBotInvalidBuyer.selector, address(0)));
        new BuyBot(owner, address(router), MIN_ORDER_AMOUNT, 0, recipient, address(0), managerRole);
    }

    function test_RevertWhen_ManagerIsZeroAddress() public {
        vm.expectRevert(abi.encodeWithSelector(BuyBot.BuyBotInvalidManager.selector, address(0)));
        new BuyBot(owner, address(router), MIN_ORDER_AMOUNT, 0, recipient, buyerRole, address(0));
    }

    function test_BuyMarketWithSufficientBalance() public {
        // Give buyer contract USDC
        uint256 buyAmount = 200e18;
        quoteToken.mint(address(buyer), buyAmount);

        // Authorized buyer can call buyMarket
        vm.prank(buyerRole);
        buyer.buyMarket(address(pair), buyAmount, 0);

        // Check that buyer contract received base tokens (since recipient is set in constructor)
        uint256 baseBalance = baseToken.balanceOf(recipient);
        assertGt(baseBalance, 0, "Should have received base tokens");

        // Check that USDC was spent
        assertEq(quoteToken.balanceOf(address(buyer)), 0, "USDC should be spent");
    }

    function test_RevertWhen_InsufficientBalance() public {
        // Give buyer contract less than minimum
        uint256 buyAmount = MIN_ORDER_AMOUNT - 1;
        quoteToken.mint(address(buyer), buyAmount);

        vm.prank(buyerRole);
        vm.expectRevert(abi.encodeWithSelector(BuyBot.BuyBotInsufficientBalance.selector, buyAmount, MIN_ORDER_AMOUNT));
        buyer.buyMarket(address(pair), buyAmount, 0);
    }

    function test_RevertWhen_AmountIsZero() public {
        quoteToken.mint(address(buyer), 200e18);

        vm.prank(buyerRole);
        vm.expectRevert(abi.encodeWithSelector(BuyBot.BuyBotInvalidAmount.selector, 0));
        buyer.buyMarket(address(pair), 0, 0);
    }

    function test_OwnerCanWithdraw() public {
        // Give buyer contract some tokens
        uint256 amount = 500e18;
        quoteToken.mint(address(buyer), amount);
        baseToken.mint(address(buyer), amount);

        uint256 ownerQuoteBalanceBefore = quoteToken.balanceOf(owner);
        uint256 ownerBaseBalanceBefore = baseToken.balanceOf(owner);

        // Withdraw USDC
        vm.startPrank(owner);
        buyer.withdraw(address(quoteToken), amount);
        assertEq(quoteToken.balanceOf(owner), ownerQuoteBalanceBefore + amount);

        // Withdraw TOKEN (0 for entire balance)
        buyer.withdraw(address(baseToken), 0);
        assertEq(baseToken.balanceOf(owner), ownerBaseBalanceBefore + amount);
        vm.stopPrank();
    }

    function test_RevertWhen_NonOwnerTriesToWithdraw() public {
        quoteToken.mint(address(buyer), 100e18);

        vm.prank(user);
        vm.expectRevert();
        buyer.withdraw(address(quoteToken), 100e18);
    }

    function test_OwnerCanSetMinOrderAmount() public {
        uint256 newMinAmount = 500e18;

        vm.prank(owner);
        vm.expectEmit(true, true, false, false);
        emit BuyBot.MinOrderAmountSet(MIN_ORDER_AMOUNT, newMinAmount);
        buyer.setMinOrderAmount(newMinAmount);

        assertEq(buyer.minOrderAmount(), newMinAmount);
    }

    function test_RevertWhen_SetMinOrderAmountToZero() public {
        vm.prank(owner);
        vm.expectRevert(abi.encodeWithSelector(BuyBot.BuyBotInvalidMinOrderAmount.selector, 0));
        buyer.setMinOrderAmount(0);
    }

    function test_CanBuyMarketView() public {
        // No balance - authorized caller
        (bool canBuy, uint256 balance) = buyer.canBuyMarket(address(pair), buyerRole);
        assertFalse(canBuy);
        assertEq(balance, 0);

        // Unauthorized caller
        (canBuy, balance) = buyer.canBuyMarket(address(pair), user);
        assertFalse(canBuy);
        assertEq(balance, 0);

        // Insufficient balance
        quoteToken.mint(address(buyer), MIN_ORDER_AMOUNT - 1);
        (canBuy, balance) = buyer.canBuyMarket(address(pair), buyerRole);
        assertFalse(canBuy);
        assertEq(balance, MIN_ORDER_AMOUNT - 1);

        // Sufficient balance
        quoteToken.mint(address(buyer), 1);
        (canBuy, balance) = buyer.canBuyMarket(address(pair), buyerRole);
        assertTrue(canBuy);
        assertEq(balance, MIN_ORDER_AMOUNT);
    }

    function test_CanBuyMarketViewWithInterval() public {
        // Deploy buyer with 60 second interval
        BuyBot buyerWithInterval =
            new BuyBot(owner, address(router), MIN_ORDER_AMOUNT, 60, recipient, buyerRole, managerRole);

        // Sufficient balance, but no lastBuyTime yet (should be true)
        quoteToken.mint(address(buyerWithInterval), MIN_ORDER_AMOUNT);
        (bool canBuy, uint256 balance) = buyerWithInterval.canBuyMarket(address(pair), buyerRole);
        assertTrue(canBuy, "Should be able to buy when no lastBuyTime");
        assertEq(balance, MIN_ORDER_AMOUNT);

        // Execute first buy
        vm.prank(buyerRole);
        buyerWithInterval.buyMarket(address(pair), MIN_ORDER_AMOUNT, 0);

        // Immediately after buy - should be false (interval not passed)
        quoteToken.mint(address(buyerWithInterval), MIN_ORDER_AMOUNT);
        (canBuy, balance) = buyerWithInterval.canBuyMarket(address(pair), buyerRole);
        assertFalse(canBuy, "Should not be able to buy immediately after");
        assertEq(balance, MIN_ORDER_AMOUNT);

        // 30 seconds later - still false
        vm.warp(block.timestamp + 30);
        (canBuy, balance) = buyerWithInterval.canBuyMarket(address(pair), buyerRole);
        assertFalse(canBuy, "Should not be able to buy after 30 seconds");

        // 60 seconds later - should be true
        vm.warp(block.timestamp + 30); // Total 60 seconds
        (canBuy, balance) = buyerWithInterval.canBuyMarket(address(pair), buyerRole);
        assertTrue(canBuy, "Should be able to buy after 60 seconds");
        assertEq(balance, MIN_ORDER_AMOUNT);
    }

    function test_GetBalance() public {
        uint256 amount = 123e18;
        quoteToken.mint(address(buyer), amount);

        assertEq(buyer.getBalance(address(quoteToken)), amount);
    }

    function test_MultipleBuys() public {
        // First buy
        uint256 firstAmount = 150e18;
        quoteToken.mint(address(buyer), firstAmount);
        vm.prank(buyerRole);
        buyer.buyMarket(address(pair), firstAmount, 0);

        uint256 firstBuyBaseBalance = baseToken.balanceOf(recipient);
        assertGt(firstBuyBaseBalance, 0);

        // Second buy
        uint256 secondAmount = 200e18;
        quoteToken.mint(address(buyer), secondAmount);
        vm.prank(buyerRole);
        buyer.buyMarket(address(pair), secondAmount, 0);

        uint256 secondBuyBaseBalance = baseToken.balanceOf(recipient);
        assertGt(secondBuyBaseBalance, firstBuyBaseBalance);
    }

    function test_BuyMarketEmitsEvent() public {
        uint256 buyAmount = 200e18;
        quoteToken.mint(address(buyer), buyAmount);

        vm.prank(buyerRole);
        vm.expectEmit(true, true, true, false);
        emit BuyBot.MarketBuyExecuted(address(pair), address(quoteToken), address(baseToken), buyAmount, buyerRole);
        buyer.buyMarket(address(pair), buyAmount, 0);
    }

    function test_BuyWithQuoteAndWithdrawBase() public {
        // Set recipient to address(0) to keep tokens in contract
        vm.prank(owner);
        buyer.setRecipient(address(0));

        // 1. Give buyer contract USDC (QUOTE)
        uint256 quoteAmount = 200e18;
        quoteToken.mint(address(buyer), quoteAmount);

        // Verify initial balances
        assertEq(quoteToken.balanceOf(address(buyer)), quoteAmount);
        assertEq(baseToken.balanceOf(address(buyer)), 0);
        assertEq(baseToken.balanceOf(owner), 0);

        // 2. Execute market buy with QUOTE tokens (keep in contract)
        vm.prank(buyerRole);
        buyer.buyMarket(address(pair), quoteAmount, 0);

        // Verify QUOTE was spent and BASE was received
        assertEq(quoteToken.balanceOf(address(buyer)), 0, "QUOTE should be fully spent");
        uint256 baseBought = baseToken.balanceOf(address(buyer));
        assertGt(baseBought, 0, "Should have received BASE tokens");

        // 3. Owner withdraws the purchased BASE tokens
        vm.prank(owner);
        buyer.withdraw(address(baseToken), 0); // 0 for entire balance

        // Verify BASE tokens were transferred to owner
        assertEq(baseToken.balanceOf(address(buyer)), 0, "BuyBot should have no BASE");
        assertEq(baseToken.balanceOf(owner), baseBought, "Owner should have all BASE");
    }

    function test_MultipleBuysAndPartialWithdraw() public {
        // Set recipient to address(0) to keep tokens in contract
        vm.prank(owner);
        buyer.setRecipient(address(0));

        // First buy
        uint256 firstAmount = 150e18;
        quoteToken.mint(address(buyer), firstAmount);
        vm.prank(buyerRole);
        buyer.buyMarket(address(pair), firstAmount, 0);
        uint256 firstBuyAmount = baseToken.balanceOf(address(buyer));

        // Second buy
        uint256 secondAmount = 200e18;
        quoteToken.mint(address(buyer), secondAmount);
        vm.prank(buyerRole);
        buyer.buyMarket(address(pair), secondAmount, 0);
        uint256 totalBought = baseToken.balanceOf(address(buyer));
        assertGt(totalBought, firstBuyAmount);

        // Partial withdraw
        uint256 withdrawAmount = totalBought / 2;
        vm.prank(owner);
        buyer.withdraw(address(baseToken), withdrawAmount);

        assertEq(baseToken.balanceOf(owner), withdrawAmount);
        assertEq(baseToken.balanceOf(address(buyer)), totalBought - withdrawAmount);
    }

    // ===== WETH (Native Coin) Tests =====

    function test_BuyWETHWithQuoteReceivesETH() public {
        // Give buyer contract USDC
        uint256 quoteAmount = 200e18;
        quoteToken.mint(address(buyerWithWETH), quoteAmount);

        // Execute market buy for WETH (will be sent to recipient)
        // WETH will automatically unwrap to ETH when sent to recipient
        vm.prank(buyerRole);
        buyerWithWETH.buyMarket(address(wethPair), quoteAmount, 0);

        // Check that USDC was spent
        assertEq(quoteToken.balanceOf(address(buyerWithWETH)), 0, "USDC should be spent");

        // Check that ETH was received by recipient (WETH auto-unwraps)
        uint256 ethBalance = recipient.balance;
        assertGt(ethBalance, 0, "Recipient should have received ETH");
    }

    function test_BuyWETHAndWithdrawETH() public {
        // Set recipient to address(0) to keep in contract
        vm.prank(owner);
        buyerWithWETH.setRecipient(address(0));

        // Give buyer contract USDC
        uint256 quoteAmount = 200e18;
        quoteToken.mint(address(buyerWithWETH), quoteAmount);

        // Execute market buy for WETH (receives ETH in contract)
        vm.prank(buyerRole);
        buyerWithWETH.buyMarket(address(wethPair), quoteAmount, 0);

        uint256 ethReceived = address(buyerWithWETH).balance;
        assertGt(ethReceived, 0, "Should have received ETH");

        // Withdraw ETH
        uint256 ownerETHBefore = owner.balance;
        vm.prank(owner);
        buyerWithWETH.withdrawETH(0);

        assertEq(owner.balance, ownerETHBefore + ethReceived);
        assertEq(address(buyerWithWETH).balance, 0);
    }

    function test_ReceiveETHAndWithdraw() public {
        // Send ETH to buyer contract
        uint256 ethAmount = 5 ether;
        vm.deal(user, ethAmount);
        vm.prank(user);
        (bool success,) = address(buyerWithWETH).call{value: ethAmount}("");
        assertTrue(success, "ETH transfer should succeed");

        // Check balance
        assertEq(address(buyerWithWETH).balance, ethAmount);
        assertEq(buyerWithWETH.getBalance(buyerWithWETH.NATIVE_COIN()), ethAmount);

        // Withdraw ETH
        uint256 ownerBalanceBefore = owner.balance;
        vm.prank(owner);
        buyerWithWETH.withdrawETH(0);

        assertEq(owner.balance, ownerBalanceBefore + ethAmount);
        assertEq(address(buyerWithWETH).balance, 0);
    }

    function test_PartialETHWithdraw() public {
        // Send ETH to buyer contract
        uint256 ethAmount = 10 ether;
        vm.deal(address(buyerWithWETH), ethAmount);

        // Partial withdraw
        uint256 withdrawAmount = 3 ether;
        uint256 ownerBalanceBefore = owner.balance;

        vm.prank(owner);
        buyerWithWETH.withdrawETH(withdrawAmount);

        assertEq(owner.balance, ownerBalanceBefore + withdrawAmount);
        assertEq(address(buyerWithWETH).balance, ethAmount - withdrawAmount);
    }

    function test_RevertWhen_NonOwnerTriesToWithdrawETH() public {
        vm.deal(address(buyerWithWETH), 5 ether);

        vm.prank(user);
        vm.expectRevert();
        buyerWithWETH.withdrawETH(1 ether);
    }

    function test_GetBalanceForETH() public {
        uint256 ethAmount = 7 ether;
        vm.deal(address(buyerWithWETH), ethAmount);

        assertEq(buyerWithWETH.getBalance(buyerWithWETH.NATIVE_COIN()), ethAmount);
    }

    // ===== New Feature Tests: Amount and Recipient Parameters =====

    function test_BuyMarketWithSpecificAmount() public {
        // Set recipient to address(0) to keep in contract
        vm.prank(owner);
        buyer.setRecipient(address(0));

        // Give buyer contract 300 USDC
        uint256 totalBalance = 300e18;
        quoteToken.mint(address(buyer), totalBalance);

        // Buy with only 150 USDC (not entire balance)
        uint256 spendAmount = 150e18;
        vm.prank(buyerRole);
        buyer.buyMarket(address(pair), spendAmount, 0);

        // Check remaining balance
        assertEq(quoteToken.balanceOf(address(buyer)), totalBalance - spendAmount, "Should have remaining QUOTE");
        assertGt(baseToken.balanceOf(address(buyer)), 0, "Should have received BASE");
    }

    function test_BuyMarketAndTransferToRecipient() public {
        // Give buyer contract USDC
        uint256 buyAmount = 200e18;
        quoteToken.mint(address(buyer), buyAmount);

        // recipient is already set in constructor
        uint256 recipientBaseBefore = baseToken.balanceOf(recipient);

        // Buy and transfer to recipient
        vm.prank(buyerRole);
        buyer.buyMarket(address(pair), buyAmount, 0);

        // Check that BASE was sent to recipient, not buyer contract
        assertEq(baseToken.balanceOf(address(buyer)), 0, "BuyBot should not have BASE");
        assertGt(baseToken.balanceOf(recipient), recipientBaseBefore, "Recipient should have received BASE");
    }

    function test_BuyMarketWithAmountAndRecipient() public {
        // Give buyer contract 300 USDC
        uint256 totalBalance = 300e18;
        quoteToken.mint(address(buyer), totalBalance);

        // recipient is already set in constructor
        uint256 spendAmount = 150e18;

        // Buy specific amount and send to recipient
        vm.prank(buyerRole);
        buyer.buyMarket(address(pair), spendAmount, 0);

        // Check remaining QUOTE in buyer
        assertEq(quoteToken.balanceOf(address(buyer)), totalBalance - spendAmount, "Should have remaining QUOTE");

        // Check BASE was sent to recipient
        assertEq(baseToken.balanceOf(address(buyer)), 0, "BuyBot should not have BASE");
        assertGt(baseToken.balanceOf(recipient), 0, "Recipient should have BASE");
    }

    function test_RevertWhen_BuyAmountExceedsBalance() public {
        // Give buyer contract 200 USDC
        uint256 balance = 200e18;
        quoteToken.mint(address(buyer), balance);

        // Try to buy with 300 USDC
        uint256 excessiveAmount = 300e18;
        vm.prank(buyerRole);
        vm.expectRevert(abi.encodeWithSelector(BuyBot.BuyBotInsufficientBalance.selector, balance, excessiveAmount));
        buyer.buyMarket(address(pair), excessiveAmount, 0);
    }

    function test_BuyWETHKeepsInContract() public {
        // Set recipient to address(0) to keep in contract
        vm.prank(owner);
        buyerWithWETH.setRecipient(address(0));

        // Give buyer contract USDC
        uint256 quoteAmount = 200e18;
        quoteToken.mint(address(buyerWithWETH), quoteAmount);

        // Buy WETH and it will auto-unwrap to ETH, stays in contract
        vm.prank(buyerRole);
        buyerWithWETH.buyMarket(address(wethPair), quoteAmount, 0);

        // ETH should be in buyer contract
        assertGt(address(buyerWithWETH).balance, 0, "BuyBot should have ETH");
    }

    // ===== Interval Tests =====

    function test_BuyMarketWithInterval() public {
        // Deploy buyer with 60 second interval
        BuyBot buyerWithInterval =
            new BuyBot(owner, address(router), MIN_ORDER_AMOUNT, 60, recipient, buyerRole, managerRole);

        // First buy should succeed
        uint256 buyAmount = 200e18;
        quoteToken.mint(address(buyerWithInterval), buyAmount);
        vm.prank(buyerRole);
        buyerWithInterval.buyMarket(address(pair), buyAmount, 0);

        // Second buy should fail (interval not passed)
        quoteToken.mint(address(buyerWithInterval), buyAmount);
        vm.prank(buyerRole);
        vm.expectRevert(abi.encodeWithSelector(BuyBot.BuyBotIntervalNotPassed.selector, 0, 60));
        buyerWithInterval.buyMarket(address(pair), buyAmount, 0);

        // Warp time forward by 59 seconds (still not enough)
        vm.warp(block.timestamp + 59);
        vm.prank(buyerRole);
        vm.expectRevert(abi.encodeWithSelector(BuyBot.BuyBotIntervalNotPassed.selector, 59, 60));
        buyerWithInterval.buyMarket(address(pair), buyAmount, 0);

        // Warp time forward by 1 more second (total 60 seconds)
        vm.warp(block.timestamp + 1);
        quoteToken.mint(address(buyerWithInterval), buyAmount);
        vm.prank(buyerRole);
        buyerWithInterval.buyMarket(address(pair), buyAmount, 0);

        // Should succeed - recipient gets tokens
        assertGt(baseToken.balanceOf(recipient), 0);
    }

    function test_OwnerCanSetInterval() public {
        uint256 newInterval = 120;

        vm.prank(owner);
        vm.expectEmit(true, true, false, false);
        emit BuyBot.IntervalSet(0, newInterval);
        buyer.setInterval(newInterval);

        assertEq(buyer.interval(), newInterval);
    }

    // ===== Authorization Tests =====

    function test_OnlyOwnerCanBuyMarket() public {
        uint256 buyAmount = 200e18;
        quoteToken.mint(address(buyer), buyAmount);

        // Owner can call buyMarket
        vm.prank(owner);
        buyer.buyMarket(address(pair), buyAmount, 0);

        assertGt(baseToken.balanceOf(recipient), 0);
    }

    function test_AuthorizedBuyerCanBuyMarket() public {
        uint256 buyAmount = 200e18;
        quoteToken.mint(address(buyer), buyAmount);

        // Authorized buyer can call buyMarket
        vm.prank(buyerRole);
        buyer.buyMarket(address(pair), buyAmount, 0);

        assertGt(baseToken.balanceOf(recipient), 0);
    }

    function test_RevertWhen_UnauthorizedCallerTriesToBuyMarket() public {
        uint256 buyAmount = 200e18;
        quoteToken.mint(address(buyer), buyAmount);

        // Unauthorized user cannot call buyMarket
        vm.prank(user);
        vm.expectRevert(abi.encodeWithSelector(BuyBot.BuyBotUnauthorizedCaller.selector, user));
        buyer.buyMarket(address(pair), buyAmount, 0);
    }

    // ===== Recipient and Buyer Setter Tests =====

    function test_OwnerCanSetRecipient() public {
        address newRecipient = makeAddr("newRecipient");

        vm.prank(owner);
        vm.expectEmit(true, true, false, false);
        emit BuyBot.RecipientSet(recipient, newRecipient);
        buyer.setRecipient(newRecipient);

        assertEq(buyer.recipient(), newRecipient);
    }

    function test_RevertWhen_NonOwnerTriesToSetRecipient() public {
        address newRecipient = makeAddr("newRecipient");

        vm.prank(user);
        vm.expectRevert();
        buyer.setRecipient(newRecipient);
    }

    function test_OwnerCanGrantBuyerRole() public {
        address newBuyer = makeAddr("newBuyer");

        vm.prank(owner);
        buyer.grantBuyerRole(newBuyer);

        assertTrue(buyer.hasRole(buyer.BUYER_ROLE(), newBuyer));

        // New buyer can now call buyMarket
        uint256 buyAmount = 200e18;
        quoteToken.mint(address(buyer), buyAmount);
        vm.prank(newBuyer);
        buyer.buyMarket(address(pair), buyAmount, 0);
    }

    function test_OwnerCanRevokeBuyerRole() public {
        vm.prank(owner);
        buyer.revokeBuyerRole(buyerRole);

        assertFalse(buyer.hasRole(buyer.BUYER_ROLE(), buyerRole));

        // Revoked buyer cannot call buyMarket
        uint256 buyAmount = 200e18;
        quoteToken.mint(address(buyer), buyAmount);
        vm.prank(buyerRole);
        vm.expectRevert(abi.encodeWithSelector(BuyBot.BuyBotUnauthorizedCaller.selector, buyerRole));
        buyer.buyMarket(address(pair), buyAmount, 0);
    }

    function test_RevertWhen_NonOwnerTriesToGrantBuyerRole() public {
        address newBuyer = makeAddr("newBuyer");

        vm.prank(user);
        vm.expectRevert();
        buyer.grantBuyerRole(newBuyer);
    }

    function test_SetRecipientToZeroKeepsTokensInContract() public {
        // Set recipient to address(0)
        vm.prank(owner);
        buyer.setRecipient(address(0));

        // Buy tokens
        uint256 buyAmount = 200e18;
        quoteToken.mint(address(buyer), buyAmount);
        vm.prank(buyerRole);
        buyer.buyMarket(address(pair), buyAmount, 0);

        // Tokens should be in contract, not recipient
        assertGt(baseToken.balanceOf(address(buyer)), 0);
        assertEq(baseToken.balanceOf(recipient), 0);
    }

    function test_IntervalZeroDisablesCheck() public {
        // Deploy buyer with 0 interval (disabled)
        BuyBot buyerNoInterval =
            new BuyBot(owner, address(router), MIN_ORDER_AMOUNT, 0, recipient, buyerRole, managerRole);

        // First buy
        uint256 buyAmount = 200e18;
        quoteToken.mint(address(buyerNoInterval), buyAmount);
        vm.prank(buyerRole);
        buyerNoInterval.buyMarket(address(pair), buyAmount, 0);

        // Second buy should succeed immediately (no interval check)
        quoteToken.mint(address(buyerNoInterval), buyAmount);
        vm.prank(buyerRole);
        buyerNoInterval.buyMarket(address(pair), buyAmount, 0);

        assertGt(baseToken.balanceOf(recipient), 0);
    }

    function test_UpdateIntervalDuringOperation() public {
        // Start with no interval
        assertEq(buyer.interval(), 0);

        // First buy
        uint256 buyAmount = 200e18;
        quoteToken.mint(address(buyer), buyAmount);
        vm.prank(buyerRole);
        buyer.buyMarket(address(pair), buyAmount, 0);

        // Set interval to 30 seconds
        vm.prank(owner);
        buyer.setInterval(30);

        // Second buy should fail (interval now active)
        quoteToken.mint(address(buyer), buyAmount);
        vm.prank(buyerRole);
        vm.expectRevert(abi.encodeWithSelector(BuyBot.BuyBotIntervalNotPassed.selector, 0, 30));
        buyer.buyMarket(address(pair), buyAmount, 0);

        // Warp time and try again
        vm.warp(block.timestamp + 30);
        quoteToken.mint(address(buyer), buyAmount);
        vm.prank(buyerRole);
        buyer.buyMarket(address(pair), buyAmount, 0);

        assertGt(baseToken.balanceOf(recipient), 0);
    }

    // ===== Manager Tests =====

    function test_ManagerCanSetMinOrderAmount() public {
        uint256 newAmount = 500e18;

        vm.prank(managerRole);
        vm.expectEmit(true, true, false, false);
        emit BuyBot.MinOrderAmountSet(MIN_ORDER_AMOUNT, newAmount);
        buyer.setMinOrderAmount(newAmount);

        assertEq(buyer.minOrderAmount(), newAmount);
    }

    function test_ManagerCanSetInterval() public {
        uint256 newInterval = 60;

        vm.prank(managerRole);
        vm.expectEmit(true, true, false, false);
        emit BuyBot.IntervalSet(0, newInterval);
        buyer.setInterval(newInterval);

        assertEq(buyer.interval(), newInterval);
    }

    function test_OwnerCanGrantManagerRole() public {
        address newManager = makeAddr("newManager");

        vm.prank(owner);
        buyer.grantManagerRole(newManager);

        assertTrue(buyer.hasRole(buyer.MANAGER_ROLE(), newManager));

        // New manager can now set minOrderAmount
        uint256 newAmount = 500e18;
        vm.prank(newManager);
        buyer.setMinOrderAmount(newAmount);
        assertEq(buyer.minOrderAmount(), newAmount);
    }

    function test_OwnerCanRevokeManagerRole() public {
        vm.prank(owner);
        buyer.revokeManagerRole(managerRole);

        assertFalse(buyer.hasRole(buyer.MANAGER_ROLE(), managerRole));

        // Revoked manager cannot set minOrderAmount
        vm.prank(managerRole);
        vm.expectRevert(abi.encodeWithSelector(BuyBot.BuyBotUnauthorizedCaller.selector, managerRole));
        buyer.setMinOrderAmount(500e18);
    }

    function test_RevertWhen_NonOwnerTriesToGrantManagerRole() public {
        address newManager = makeAddr("newManager");

        vm.prank(user);
        vm.expectRevert();
        buyer.grantManagerRole(newManager);
    }

    function test_RevertWhen_UnauthorizedUserTriesToSetMinOrderAmount() public {
        vm.prank(user);
        vm.expectRevert(abi.encodeWithSelector(BuyBot.BuyBotUnauthorizedCaller.selector, user));
        buyer.setMinOrderAmount(500e18);
    }

    function test_RevertWhen_UnauthorizedUserTriesToSetInterval() public {
        vm.prank(user);
        vm.expectRevert(abi.encodeWithSelector(BuyBot.BuyBotUnauthorizedCaller.selector, user));
        buyer.setInterval(60);
    }
}
