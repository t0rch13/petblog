$(document).ready(function () {
    const pets = $('#pets');

    $.ajax({
        url: 'http://localhost:8000/api/pets/',
        method: 'GET',
        success: function (response) {
            let html='';
            response.forEach(pet => {
                html += `
                <div class="col-md-4 p-4">
                    <div class="card" style="width: 18rem;">
                        <div class="card-body">
                        <h5 class="card-title">${pet.petname}</h5>
                        <h6 class="card-subtitle mb-2 text-muted">${pet.species}</h6>
                        <p class="card-text">${pet.age} years old</p>
                        </div>
                    </div>
                </div>
                `;
            });

            pets.prepend(html);
        },
        error: function (error) {
            console.log(error);
        }
    });
});