// pages/talksheet/talksheet.js
var app = getApp();
var util = require('../../utils/util.js');
var api = require('../../config/api.js');

Page({

  /**
   * 页面的初始数据
   */
  data: {
    comments: [],
    typing: '',
    input: '',
  },
  inputChange: function(e) {
    this.setData({
      input: e.detail.value
    })
  },
  sendMsg: function() {
    let that = this
    var input = this.data.input
    this.setData({
      typing: '',
      input: '',
    })

    if (input.trim() == '') {
      return
    }
    util.request(api.SheetsMsg, {
      message: input,
    })

    that.onPullDownRefresh()
  },
  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    let that = this;
    util.request(api.SheetsIndex).then(function(res) {
      if (res.errno === 0) {
        that.setData({
          comments: res.data,
        })
      }
    })
  },
  /**
   * 生命周期函数--监听页面初次渲染完成
   */
  onReady: function () {

  },

  /**
   * 生命周期函数--监听页面显示
   */
  onShow: function () {

  },

  /**
   * 生命周期函数--监听页面隐藏
   */
  onHide: function () {

  },

  /**
   * 生命周期函数--监听页面卸载
   */
  onUnload: function () {

  },

  /**
   * 页面相关事件处理函数--监听用户下拉动作
   */
  onPullDownRefresh: function () {
    this.onLoad();
  },

  /**
   * 页面上拉触底事件的处理函数
   */
  onReachBottom: function () {

  },

  /**
   * 用户点击右上角分享
   */
  onShareAppMessage: function () {

  }
})