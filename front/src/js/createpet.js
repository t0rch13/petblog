$(document).ready(function () {
    $('#create-btn').on('click', function(event) {
        const isFormValid = $('#create-form')[0].checkValidity();
    
        if(isFormValid){
            event.preventDefault();
            const petname = $('#petname').val();
            const species = $('#species').val();
            const breed = $('#breed').val();
            const age = parseInt($('#age').val(), 10);

            const data = {
                petname : petname,
                species : species,
                breed : breed,
                age : age
            }

            $.ajax({
                url: 'http://localhost:8000/api/pets/create/',
                method: 'POST',
                data: JSON.stringify(data),
                contentType: 'application/json',
                headers: {
                    'Authorization': 'Bearer ' + localStorage.getItem('authToken')
                },
                success: function (response) {
                    alert(response.petname + ' is successfully added!');
                    window.location.href = '/pets';
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