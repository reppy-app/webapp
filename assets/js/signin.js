$('#signin-form').on('submit', signin);

function signin(event) {
    event.preventDefault();

    $.ajax({
        url: "/login",
        method: "POST",
        data: {
            email: $('#inputEmail').val(),
            password: $('#inputPassword').val()
        }
    }).done(function () {
        window.location = "/home";
    }).fail(function () {
        alert("invalid email or password")
    });
}