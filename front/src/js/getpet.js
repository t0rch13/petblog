$(document).ready(function () {
    const petId = window.location.pathname.split('/')[2]; // Get the pet ID from the URL

    $.ajax({
        url: `http://localhost:8000/api/pets/${petId}`,
        method: 'GET',
        headers: {
            'Authorization': 'Bearer ' + localStorage.getItem('authToken')
        },
        success: function (response) {
            console.log(response);
            const pet = response;
            console.log(pet);
            $('#pet-name').text(pet.petname);
            $('#pet-species').text(pet.species);
            $('#pet-breed').text(pet.breed);
            $('#pet-age').text(pet.age);
        },
        error: function (error) {
            console.log(error);
        }
    });
});