$(document).ready(function () {
    $.ajax({
        url: 'http://localhost:8000/api/pets/', 
        method: 'GET',
        headers: {
            'Authorization': 'Bearer ' + localStorage.getItem('authToken')
        },
        success: function (response) {
            // Populate the pet select field with the user's pets
            const petsArr = response.data;
            const petSelect = $('#pet');
            petsArr.forEach(pet => {
                petSelect.append(new Option(pet.petname + ' ' + pet.species , pet.id));
                console.log(pet);
            });
        },
        error: function (error) {
            console.log(error);
        }
    });

    $('#create-btn').on('click', function(event) {
        const isFormValid = $('#create-form')[0].checkValidity();
    
        if(isFormValid){
            event.preventDefault();
            const title = $('#title').val();
            const content = $('#content').val();
            const petId = $('#pet').val();

            console.log(petId);
    
            const data = {
                title : title,
                content : content
            }
    
            $.ajax({
                url: 'http://localhost:8000/api/posts/create/' + petId, // Append the pet's id to the URL
                method: 'POST',
                data: JSON.stringify(data),
                contentType: "application/json",
                headers: {
                    'Authorization': 'Bearer ' + localStorage.getItem('authToken')
                },
                success: function (response) {
                    window.location.href = '/';
                },
                error: function (error) {
                    console.log(error);
                }
            });
    
        }else{
            $('#create-form')[0].reportValidity();
        }
    })
});