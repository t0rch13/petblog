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

    $('#editPetModal').on('show.bs.modal', function (event) {
        var button = $(event.relatedTarget) 
        var petName = $('#pet-name').text() 
        var petSpecies = $('#pet-species').text()
        var petBreed = $('#pet-breed').text()
        var petAge = $('#pet-age').text()
      
        // Update the modal's content.
        var modal = $(this)
        modal.find('#pet-name-input').val(petName)
        modal.find('#pet-species-input').val(petSpecies)
        modal.find('#pet-breed-input').val(petBreed)
        modal.find('#pet-age-input').val(petAge)
    })
      
      
    $('#saveChangesButton').click(function() {
        var petName = $('#pet-name-input').val()
        var petSpecies = $('#pet-species-input').val()
        var petBreed = $('#pet-breed-input').val()
        var petAge = parseInt($('#pet-age-input').val()) // Parse age as integer
    
        var data = {
            petname: petName,
            species: petSpecies,
            breed: petBreed,
            age: petAge
        };
    
        $.ajax({
            url: `http://localhost:8000/api/pets/${petId}`,
            method: 'PUT',
            headers: {
                'Authorization': 'Bearer ' + localStorage.getItem('authToken')
            },
            data: JSON.stringify(data), 
            contentType: 'application/json', 
            success: function(response) {
                console.log(response);
                $('#editPetModal').modal('hide');
                $('#pet-name').text(petName);
                $('#pet-species').text(petSpecies);
                $('#pet-breed').text(petBreed);
                $('#pet-age').text(petAge);
            },
            error: function(error) {
                console.log(error);
            }
        });
    });

    $('#deletePetButton').click(function() {
        var confirmDelete = confirm("Are you sure you want to delete this pet?");
        if (confirmDelete) {
            $.ajax({
                url: `http://localhost:8000/api/pets/${petId}`,
                method: 'DELETE',
                headers: {
                    'Authorization': 'Bearer ' + localStorage.getItem('authToken')
                },
                success: function(response) {
                    console.log(response);
                    window.location.href = '/pets';
                },
                error: function(error) {
                    console.log(error);
                }
            });
        }
    });
});