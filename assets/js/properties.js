$('#create-reppy').on('submit', createReppy)

function createReppy(event) {
    event.preventDefault();

    $.ajax({
        url: "/properties",
        method: "POST",
        data: {
            name: $('#name').val(),
            price: $('#price').val()
        }
    }).done(function () {
        window.location = "/home";
    }).fail(function () {
        alert("Erro ao criar república!");
    });
}