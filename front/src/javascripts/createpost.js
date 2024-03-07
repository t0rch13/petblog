$(document).ready(function () {
    $('#create-btn').on('click', function(event) {
        const isFormValid = $('#create-form')[0].checkValidity();
    
        if(isFormValid){
            event.preventDefault();
            const petname = $('#petname').val();
            const title = $('#title').val();
            const content = $('#content').val();
            const pet = $('#pet').val();

            const data = {
                petname : petname,
                title : title,
                content : content
            }

            $.ajax({
                url: 'http://localhost:8000/api/posts/create/',
                method: 'POST',
                data : data,
                success: function (response) {
                    
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