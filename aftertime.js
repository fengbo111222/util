async function a() {
  await sleep(1000)
  console.log(`-----=====`)


}
a()
async function sleep(ms) {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

module.exports.utilSleep = sleep