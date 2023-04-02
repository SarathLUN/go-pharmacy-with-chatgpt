$(document).ready(function () {
    $('#btn-submit-order').click(function () {
        let apiToken = "5435733281:AAGIQWMpePWTkh_5LAtoeanxhlkO4B9XyQ4";
        // share contact with @tonyStark_test_bot, then use Postman to get Telegram user_id
        let chatId = "508976051";
        let text = 'Hello, You got an order: [123](https://google.com)'; // link back to /pharmacy-html/shop-cart.html

        let urlString = `https://api.telegram.org/bot${apiToken}/sendMessage?chat_id=${chatId}&text=${text}&parse_mode=Markdown`;

        let request = new XMLHttpRequest();
        request.open("GET", urlString);
        request.send();
    });
});