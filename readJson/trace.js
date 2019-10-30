
const simpleTool = require(`./simpleTool`).getJsonFromConfig;

// 所有设备位置信息列表
module.exports.getAllDevicePositionInfoList = cFun.awaitHandlerFactory(async (req, res) => {


  var deviceType = await simpleTool();

  const devicePositionList = await bIODal.getAllDevicePositionInfoList(

    deviceType[`deviceType`],

  );
});
