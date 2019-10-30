const fs = require('fs');
const path = require('path');

module.exports.getJsonFromConfig = async function() {
  var promise = new Promise((resolve, reject) => {
    fs.readFile(path.join(__dirname, '../config/config.json'), `utf-8`, (err, res) => {
      if (err) {
        reject(err);
      }
      resolve(JSON.parse(res));
    });
  });
  var data = await promise;
  console.log(`-----data=====`, data);
  return data;
};
