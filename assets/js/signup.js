$('#signup-form').on('submit', createUser);

function createUser(event) {
    event.preventDefault();
    
    if ($('#inputPassword').val() != $('#inputConfirmPassword').val()) {
        alert("Password are different");
        return;
    }

    $.ajax({
        url: "/users",
        method: "POST",
        data: {
            email: $('#inputEmail').val(),
            password: $('#inputPassword').val()
        }
    })
}