$('#modal').on('show.bs.modal', function(event) {
    var button = $(event.relatedTarget)
    var appeal = button.data('whatever')
    console.log(appeal)
    $('#notice').prop("data", "/notice?appeal="+appeal)
})