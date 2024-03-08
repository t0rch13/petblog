$(document).ready(function () {
    $('#login-btn').on('click', function (event) {
        let isFormValid = $('#login-form')[0].checkValidity();
        if (isFormValid) {
            event.preventDefault();
            const username = $('#username').val();
            const password = $('#password').val();

            const data = {
                Username: username,
                Password: password
            }

            $.ajax({
                url: 'http://localhost:8000/auth/sign-in/',
                method: 'POST',
                data: JSON.stringify(data),
                contentType: "application/json; charset=utf-8",
                dataType: "json",
                success: function (response) {
                    alert(response.token);
                    localStorage.setItem('authToken', response.token);
                    localStorage.setItem('username', username);
                    window.location.href = '/';
                },
                error: function (error) {
                    console.log(error);
                }
            });
        } else {
            $('#login-form')[0].reportValidity();
        }
    });
});