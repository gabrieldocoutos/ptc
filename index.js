import puppeteer from 'puppeteer';
// Or import puppeteer from 'puppeteer-core';

// Launch the browser and open a new blank page
const browser = await puppeteer.launch({   args: ['--disable-features=HttpsFirstBalancedModeAutoEnable'], });

const page = await browser.newPage();

// Navigate the page to a URL.
await page.goto('http://www.parquetenisclube.com.br/Booking/Grid.aspx');

// Set screen size.
await page.setViewport({width: 1080, height: 1024});





console.log("foo")


await page.waitForSelector("#tituloColumna")

await page.screenshot({
  path: 'hn.png',
});


const data = await page.evaluate(() => window.dataTable.Columnas)


console.log(data)




await browser.close();
