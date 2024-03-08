$(document).ready(function () {
    const pets = $('#pets');

    $.ajax({
        url: 'http://localhost:8000/api/pets/',
        method: 'GET',
        headers: {
            'Authorization': 'Bearer ' + localStorage.getItem('authToken')
        },
        success: function (response) {
            const petsArr = response.data;
            if (petsArr.length === 0) {
                pets.prepend('<h3>No pets added yet.</h3>');
            } else {
                let html='';
                petsArr.forEach(pet => {
                    html += `
                    <div class="col-md-4 p-4">
                        <div class="card" style="width: 18rem;">
                            <div class="card-body">
                            <h5 class="card-title">${pet.petname}</h5>
                            <h6 class="card-subtitle mb-2 text-muted">${pet.species}</h6>
                            <button onclick="window.location.href='/pet/${pet.id}'" class="btn btn-primary">View</button>
                            </div>
                        </div>
                    </div>
                    `;
                });

                pets.prepend(html);
            }
        },
        error: function (error) {
            console.log(error);
        }
    });

    $('#add-pet-btn').on('click', function () {
        window.location.href = '/createpet';
    });
});