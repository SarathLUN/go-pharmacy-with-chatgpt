<!-- script get current GeoLocation -->
$(document).ready(function () {
    $("#location-basic-addon2").click(function () {
        console.log(GeolocationPosition);
        if (GeolocationPosition) {
            navigator.geolocation.getCurrentPosition(locationSuccess);
        } else {
            $("#locationData").val('Your browser does not support location data retrieval.');
        }
    });

    function locationSuccess(position) {
        console.log(position);
        let latitude = position.coords.latitude;
        let longitude = position.coords.longitude;
        $("#locationData").val("Latitude:" + latitude + "; Longitude:" + longitude);

        // syntax:   https://maps.google.com?q=latitude,longitude
        // example:   https://maps.google.com?q=11.556305,104.921831
    }
});
