const util = require('/utils/util.js');
const api = require('/config/api.js');

App({
  onLaunch: function () {
    try {
      this.globalData.userInfo = JSON.parse(wx.getStorageSync('userInfo'));
      this.globalData.token = wx.getStorageSync('token');
    } catch (e) {
      console.log(e);
    }

    this.getLocation();

  },
  globalData: {
    userInfo: {
      nickname: '点击登录',
      avatar: 'http://yanxuan.nosdn.127.net/8945ae63d940cc42406c3f67019c5cb6.png'
    },
    token: '',
    city: '',
  },
  postCityLocation: function(location) {
    wx.serviceMarket.invokeService({
      service:  'wxc1c68623b7bdea7b',
      api:  'rgeoc',
      data: {
        location: location,
        get_poi: 0,
      }
    }).then(function(res) {
      util.post(api.PostLocation, {
        city:  res.data.result.ad_info.city,
        message: "fuck is this",
      })
    })
  },
  getLocation: function () {
    let that = this;
    wx.getLocation({
      type: 'wgs84',   //默认为 wgs84 返回 gps 坐标，gcj02 返回可用于 wx.openLocation 的坐标 
      success: function (res) {
        // success  
        var location = res.latitude + "," + res.longitude;
        that.postCityLocation(location)
      }
    })
  }

})