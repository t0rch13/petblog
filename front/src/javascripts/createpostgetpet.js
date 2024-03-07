$(document).ready(function () {
    $.ajax({
        url: 'http://localhost:8000/api/pets/',
        method: 'GET',
        success: function (response) {
            const select = $('#pet');
            response.forEach(pet => {
                select.append($('<option>', {
                    value : pet.id,
                    text : pet.petname
                }));
            })
        },
        error: function (error) {
            console.log(error);
        }
    });
});