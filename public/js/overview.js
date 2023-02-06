

async function register(userID) {
    var bdy = {
        "UserID": userID, 
    }

    fetch("/api/coffeedate/register", {
        "method": "POST",
        "body": JSON.stringify(bdy),
        "headers": {
            "Content-Type": "application/json"
        },
        "credentials": "include",
        "redirect": "follow"
    }).then(function (response) {
        return response.json();
    }).then(function (data) {
        console.log("received")
        console.log(data)
        location.reload()
        return Promise.resolve(data)
    }).catch(err => {
        if (err.name === "AbortError") {
            return
        }
    })

}


async function activate(id) {
    fetch(`/api/coffeedate/${id}/activate`, {
        "method": "POST",
        "headers": {
            "Content-Type": "application/json"
        },
        "credentials": "include",
        "redirect": "follow"
    }).then(function (response) {
        return response.json();
    }).then(function (data) {
        console.log("received")
        console.log(data)
        location.reload()
        return Promise.resolve(data)
    }).catch(err => {
        if (err.name === "AbortError") {
            return
        }
    })

}


async function deactivate(id) {
    fetch(`/api/coffeedate/${id}/deactivate`, {
        "method": "POST",
        "headers": {
            "Content-Type": "application/json"
        },
        "credentials": "include",
        "redirect": "follow"
    }).then(function (response) {
        return response.json();
    }).then(function (data) {
        console.log("received")
        console.log(data)
        location.reload()
        return Promise.resolve(data)
    }).catch(err => {
        if (err.name === "AbortError") {
            return
        }
    })

}

document.addEventListener('DOMContentLoaded', function () {
    document.getElementById('book').addEventListener('click', event => {
        console.log("will book")
        var button = event.target
        console.log(button.dataset)
        var userID = button.dataset.userId
        console.log(userID)
        register(userID)
    })


    var elem = document.querySelectorAll('.deactivateReg')
    if (elem != null) {
        elem.forEach(item => {
            item.addEventListener('click', event => {
                var btn = event.target.parentNode
                console.log(btn)
                var id = btn.dataset.id 
                console.log(id)
                deactivate(id)
            })
        })
    }

    var elem = document.querySelectorAll('.activateReg')
    if (elem != null) {
        elem.forEach(item => {
            item.addEventListener('click', event => {
                var btn = event.target.parentNode
                var id = btn.dataset.id 
                console.log(btn)
                console.log(id)
                activate(id)
            })
        })
    }

})