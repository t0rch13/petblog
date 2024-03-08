$(document).ready(function () {
    var username = localStorage.getItem('username');
    if (username) {
        $('#navbarDropdown').text(username);
    }

    $('#logoutButton').click(function() {
        localStorage.clear();
        location.href = '/'; 
    });
});