//nodejs程序，随机生成几只股票并且随机时间更新行情数据

var net = require('net');

var stknum = process.argv[2] || 5;  //默认生成五只股票

var stocks = [];

//随机生成五个不重复的股票代码
function getStkcodes(num) {
	if(num < 0)
		return;

	num *= 2;    //扩大范围，如果随机五个数就在10个里选五个
	let randomArray = [];
	let resultArray = [];

	for(let i = 0; i < num; i++){
		randomArray[i] = i;
	}

	for(let i = 0; i< num; i++){
		let seed = Math.floor(Math.random()*(num - i));    //从剩下的随机数里生成
		resultArray[i] = randomArray[seed];    //赋值给结果数组 
		randomArray[seed] = randomArray[num - i - 1];    //把随机数产生过的位置替换为未被选中的值。
	}

	return resultArray;
}

//初始化证券信息（模拟昨日收盘价，便于判断是否涨跌停）
getStkcodes(stknum).forEach((item, i) => {
	if(i >= stknum){
		return;
	}
	else{
		let stock ={};
		stock.stkcode = 600000 + item;
		stock.stkprice = Math.floor(10 + Math.random() * 10);
		stock.amount = 1000 + Math.floor(100 * (Math.random()* 10));
		//stock.buy = stock.sell;
		//stock.amount = stock.sell;
		stocks.push(stock);
	}
});

var latestStocks = stocks;

function upedateStocks() {
	latestStocks.forEach((item, i) => {
		item.stkprice = stocks[i].stkprice * (1 + Math.floor(Math.random() * 10) / 100) 
										   * (1 - Math.floor(Math.random() * 10) / 100);  //涨跌幅不超过10%
		item.stkprice = item.stkprice.toFixed(2);  //保留两位小数

		item.amount = item.amount + Math.floor(100 * (Math.random()* 10));
		//item.buy = item.sell;
		//item.amount += item.sell;
	});
}

//console.log(stocks);

var client = net.connect({host: 'localhost', port: 1377},() => {
	console.log('Connect to server.');
	setInterval(() => {
		var rsec = 3 + Math.floor(Math.random()*10)
		setTimeout(() => {
			upedateStocks();
			console.log(latestStocks);
			client.write(JSON.stringify(latestStocks));
		}, rsec);		
	}, 5000);
});

client.on('data', function(data){
	console.log('Other user\'s input ', data.toString());
});

client.on('end', function(){
	console.log('Disconnect from server');
	process.exit();
});
