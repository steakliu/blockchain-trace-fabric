import Vue from 'vue'
import App from './App'

Vue.config.productionTip = false
App.mpType = 'app'
var Fly=require("flyio/dist/npm/wx") 
var fly=new Fly
Vue.prototype.$httpBlock = fly
Vue.prototype.$httpUrl = "http://116.198.160.39:8080/route";

const app = new Vue(App)
app.$mount()
