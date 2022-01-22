function IniciarSesion(){
    var username = document.querySelector('#usuario').value
    var contraseña = document.querySelector('#contraseña').value

    var credenciales = {
        'usuario': username,
        'contraseña': contraseña
    }
    console.log(credenciales)
    fetch('http://127.0.0.1:3000/admin', {
        method: 'POST',
        body: JSON.stringify(credenciales),
        headers: {
            'Content-Type': 'application/json'
        }
    }).then(res => res.json())
        .catch(err => {
            console.error('Error:', err)
            alert("Ha ocurrido un error durante la ejecución")
        })
        .then(response => {
            console.log(response);
            console.log(response.message)
            if (response.message == "Failed") {
                alert("No has ingresado las credenciales correctas")
            } else {
                if (response.tipoUs == "1"){
                    alert(`Credenciales correctas, bienvenido ${username}. Haz ingresado como administrador.`)
                    sessionStorage.setItem("usuario", username)
                    location.href = "{{ url_for('inicioAdmin') }}"
                } else{
                    alert(`Credenciales correctas, bienvenido ${username}. Haz ingresado como administrador.`)
                    sessionStorage.setItem("usuario", username)
                    location.href = "{{ url_for('inicioAdmin') }}"
                }
            }
        })
}