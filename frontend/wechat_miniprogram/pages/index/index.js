const util = require('../../utils/util.js');
const api = require('../../config/api.js');
const user = require('../../services/user.js');

//获取应用实例
const app = getApp()
Page({
  data: {
    postGoods: [],
    scrollTop: 0,
    scrollHeight: 0,
    page: 1,
    size: 10000,
    id: 1
  },
  onShareAppMessage: function () {
    return {
      title: 'ExStuff',
      desc: '值得信赖的个人二手交易平台',
      path: '/pages/index/index'
    }
  },

  getIndexData: function () {
    let that = this;
    util.request(api.IndexUrl).then(function (res) {
      if (res.errno === 0) {
        that.setData({
          postGoods: res.data.postGoods,
        });
      }
    });

  },
  onLoad: function (options) {
    this.getIndexData();
  },
  onReady: function () {
    // 页面渲染完成
  },
  onShow: function () {
    // 页面显示
  },
  onHide: function () {
    // 页面隐藏
  },
  onUnload: function () {
    // 页面关闭
  },
})
