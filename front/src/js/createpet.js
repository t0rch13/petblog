$(document).ready(function () {
    $('#create-btn').on('click', function(event) {
        const isFormValid = $('#create-form')[0].checkValidity();
    
        if(isFormValid){
            event.preventDefault();
            const petname = $('#petname').val();
            const species = $('#species').val();
            const breed = $('#breed').val();
            const age = $('#age').val();

            const data = {
                petname : petname,
                species : species,
                breed : breed,
                age :age
            }

            $.ajax({
                url: 'http://localhost:8000/api/pets/create/',
                method: 'POST',
                data : data,
                success: function (response) {
                    const pets = $('#pets');
                    var html = `
                    <div class="col-md-4 p-4">
                        <div class="card" style="width: 18rem;">
                            <div class="card-body">
                            <h5 class="card-title">${response.petname}</h5>
                            <h6 class="card-subtitle mb-2 text-muted">${response.species}</h6>
                            <p class="card-text">${response.age} years old</p>
                            </div>
                        </div>
                    </div>
                    `;

                    pets.prepend(html);

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