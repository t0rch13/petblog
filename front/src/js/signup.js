$(document).ready(function () {
    $('#login-btn').on('click', function (event) {
        let isFormValid = $('#login-form')[0].checkValidity();
        if (isFormValid) {
            event.preventDefault();
            const name = $('#name').val();
            const username = $('#username').val();
            const password = $('#password').val();

            const data = {
                Name : name,
                Username: username,
                Password: password
            }

            $.ajax({
                url: 'http://localhost:8000/auth/sign-up/',
                method: 'POST',
                data: JSON.stringify(data),
                contentType: "application/json; charset=utf-8",
                success: function (response) {
                    alert('Signup successful!');
                    window.location.href = '/login';
                },
                error: function (error) {
                    alert('Signup failed: ' + error.responseText);
                }
            });
        } else {
            $('#login-form')[0].reportValidity();
        }
    });
});