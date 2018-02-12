$('#appeals tr').click(function() {
    var href = $(this).data("href")
    console.log("Test")
    window.location = href
})