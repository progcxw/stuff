var util = require('../../../utils/util.js');
var api = require('../../../config/api.js');
var websocket = require('../../../services/websocket.js');
// 参考:
// https://blog.csdn.net/qq_35713752/article/details/78688311
// https://blog.csdn.net/qq_35713752/article/details/80811397
Page({

  /**
   * 页面的初始数据
   */
  data: {
    id: 0,
    history_list: [],
    other_side: {},
    goods: {},
    is_u1: false,
    my_avatar: '',
    scroll_top: 0,
    offset_time: null,
    size: 10,
    scroll_height: 0,
    no_more: false,
    input: '',
    typing: '',
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function(options) {
    let now = new Date();
    this.setData({
      id: Number(options.id),
      my_avatar: wx.getStorageSync('userInfo').avatar,
      offset_time: now.toISOString()
    })

    this.getHistory();
    this.openListen();

  },

  getHistory: function() {
    let that = this;
    util.request(api.ChatForm + '?chat_id=' + this.data.id, {
      offset_time: this.data.offset_time,
      size: this.data.size
    }).then(function(res) {
      if (res.errno === 0) {
        console.log(res.data);
        
        that.setData({
          other_side: res.data.other_side,
          history_list: res.data.history_list.concat(that.data.history_list),
          goods: res.data.goods,
          is_u1: res.data.is_u1,
          offset_time: res.data.offset_time+"",
        });
        
        if (res.data.history_list.length < that.data.size) {
          that.setData({
            no_more: true
          })
        }
        if (that.data.history_list.length < 11) {
          //首次加载
          wx.setNavigationBarTitle({
            title: that.data.other_side.nick_name
          })

          let _this = that
          that.getScrollHeight().then((res) => {
            var scroll = res - _this.data.scroll_height
            _this.setData({
              scroll_top: 5000,
              scroll_height: res,
            })
          })

        } else {
          //重新设置scroll,scroll_top = 之前的scrollHeigth - 加入了数据后的scrollHeigth
          let _this = that
          that.getScrollHeight().then((res) => {
            var scroll = res - _this.data.scroll_height
            _this.setData({
              scroll_top: scroll,
              scroll_height: res,
            })
          })

        }
      } else {
        console.log(res)
      }
    })
  },
  openListen:function(){
    let that = this
    websocket.listenChatForm(this.data.id).then(res => {
      var newHistory = [{
        chat_id: res.chat_id,
        u1_to_u2: res.senderId < res.receiverId ? true : false,
        message_type: res.message_type,
        message_body: res.message_body,
        send_time: res.send_time,
      }]
      that.addHistoryList(newHistory)
      that.openListen()
    })
  },
  toGoods: function(event) {
    let goodsId = event.target.dataset.id;
    wx.navigateTo({
      url: '/pages/postgoods/postgoods?id=' + goodsId,
    });
  },
  more: function() {
    console.log("到顶加载更多")
    if (!this.data.no_more) {
      this.getHistory()
    }
  },
  getScrollHeight: function() {
    let that = this
    return new Promise(function(resolve, reject) {
      var query = wx.createSelectorQuery()
      //#hei是位于scroll最低端的元素,求它离scroll顶部的距离,得出ScrollHeight
      query.select('#hei').boundingClientRect()
      query.selectViewport().scrollOffset()
      query.exec(function(res) {
        console.log("Scroll_height " + res[0].top)
        resolve(res[0].top);
      })
    });
  },
  inputChange: function(e) {
    console.log(e.detail.value)
    this.setData({
      input: e.detail.value
    })
  },
  sendMsg: function() {

    let that = this
    var data = this.createMsg()
    var input = this.data.input
    this.setData({
      typing: '',
      input: '',
    })
    if (input.trim() == '') {
      return
    }
    //通过webSocket发送消息
    websocket.sendMessage(data).then(res => {
      console.log(res)
      console.log(this.data.other_side.user_id)

      var newHistory = [{
        chat_id: this.data.id,
        u1_to_u2: wx.getStorageSync('userInfo').id > this.data.other_side.user_id ? true : false,
        message_type: 1,
        message_body: input,
        send_time: util.formatTime(new Date()),
      }]

      that.addHistoryList(newHistory)


    }).catch((res) => {
      console.log(res)
      util.showErrorToast('发送失败')
    });

  },
  addHistoryList: function(history_list) {
    //把新的数据加入目前的对话框
    var newHistoryList = this.data.history_list.concat(history_list)
    this.setData({
      history_list: newHistoryList,
    })


    //重新设置scroll
    let _this = this
    this.getScrollHeight().then((res) => {
      var scroll = res - _this.data.scroll_height
      _this.setData({
        scroll_top: 100000000,
        scroll_height: res,
      })
    })
  },
  createMsg: function() {
    var msgType;
  
    if (this.data.history_list.length>1) {
      msgType = 1
    } else {
      msgType = 3
    }

    var data = JSON.stringify({
      chat_id: this.data.id,
      sender_id: wx.getStorageSync('userInfo').id,
      receiver_id: this.data.other_side.user_id,
      goods_id: this.data.goods.id,
      message_type: msgType,
      message_body: this.data.input,
      send_time: 0
    })

    console.log("message: ", data)
    return data
  },
  buy:function(){
    util.showErrorToast('功能开发中')
  },

  /**
   * 生命周期函数--监听页面初次渲染完成
   */
  onReady: function() {

  },

  /**
   * 生命周期函数--监听页面显示
   */
  onShow: function() {

  },

  /**
   * 生命周期函数--监听页面隐藏
   */
  onHide: function() {
    websocket.listenBadge()
  },

  /**
   * 生命周期函数--监听页面卸载
   */
  onUnload: function() {
    
  },

  /**
   * 页面相关事件处理函数--监听用户下拉动作
   */
  onPullDownRefresh: function() {

  },

  /**
   * 页面上拉触底事件的处理函数
   */
  onReachBottom: function() {

  },

  /**
   * 用户点击右上角分享
   */
  onShareAppMessage: function() {

  },


})