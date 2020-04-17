const fs = require('fs');
const path = require('path');
const simpleTool = require(`./simpleTool`);

//读写需要同步方法，因为本来进入时候就异步了
module.exports.getJsonFromConfig = async function (fileName) {

  var promise = new Promise(async function (resolve, reject) {
    // fs.readFile(path.join(__dirname, '../config/' + fileName), `utf-8`, async function (err, res) {
    //   if (err) {
    //     reject(err);
    //   }
    //   let resData = await JSON.parse(res)
    //   resolve(resData);

    // });

    let res = fs.readFileSync(path.join(__dirname, '../config/' + fileName), `utf-8`)
    let resData = await JSON.parse(res);
    resolve(resData);
  });



  var data = await promise;
  // console.log(`-----data=====`, data);
  return data;
};
module.exports.changeJsonFromConfig = async function (fileName, time, beginName) {

  try {
    let modelInfo = await simpleTool.getJsonFromConfig(fileName)

    modelInfo[beginName] = time
    var str = await JSON.stringify(modelInfo);
    // await fs.writeFile(path.join(__dirname, '../config/' + fileName), str, async function (err) {
    //   if (err) {
    //     console.error(err);
    //   }
    //   let modelInfo = await simpleTool.getJsonFromConfig(fileName)

    //   console.log('--------------------修改modelInfo-deviceAlarmInsertTime成功', modelInfo.beginTime);
    // })
    await fs.writeFileSync(path.join(__dirname, '../config/' + fileName), str)
    modelInfo = await simpleTool.getJsonFromConfig(fileName)

    console.log('--------------------修改modelInfo-deviceAlarmInsertTime成功----', beginName + "--" + modelInfo[beginName]);

  } catch (error) {
    console.log(`-----=====`, error)

  }
}



// -----------------
// 从数据读出所需并写入json文件
var e_sense_alarm_alarmType = require('../../dal/e_sense_alarm_alarmType');
const fs = require('fs');
const path = require('path');

// 每四分钟循环一次的任务
module.exports.getFlowStepByInsterTime = async (modeType, modelCode) => {

  let info = await e_sense_alarm_alarmType.getAlarmInfo()
  console.log(`-----=====`, JSON.stringify(info))

  for (let [_, elem] of info.entries()) {

    elem["timeRule"] = "* * * * * *"

  }
  data = {
    "model": info
  }

  // 将当前时间插入json表

  await fs.writeFileSync(path.join(__dirname, '../../../config/' + "modelInfo.json"), JSON.stringify(data))
  return
}
