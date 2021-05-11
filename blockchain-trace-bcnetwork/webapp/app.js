const express = require("express");
const app = new express();

const bodyParser = require('body-parser');

//(3)使用中间件 : 给req添加body属性，存储解析好的post参数
app.use(bodyParser.urlencoded({ extended: false }));
app.use(bodyParser.json());

app.all("*",function(req,res,next){
    //设置允许跨域的域名，*代表允许任意域名跨域
    res.header("Access-Control-Allow-Origin","*");
    //允许的header类型
    res.header("Access-Control-Allow-Headers","content-type");
    //跨域允许的请求方式   
    res.header("Access-Control-Allow-Methods","DELETE,PUT,POST,GET,OPTIONS");
    if (req.method.toLowerCase() === 'options')
        res.send(200);  //让options尝试请求快速结束
    else
        next();
});      

//引入route模块
const route  = require("./route/route");

const port = 8080

//加载route模块
app.use("/route",route);

app.listen(port, () => console.log(`Example app listening on port ${port}!`))
