document.getElementById("login-form").addEventListener("submit", async function(event) {
    event.preventDefault();

    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;

    const response = await fetch("/login-action", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({ username, password })
    });

    const data = await response.json();
    if (response.ok) {
        localStorage.setItem("token", data.token);
        window.location.href = "/protected";
    } else {
        alert(data.error);
    }
});

document.getElementById("register-form").addEventListener("submit", async function(event) {
    event.preventDefault();

    const username = document.getElementById("username").value;
    const password = document.getElementById("password").value;

    const response = await fetch("/register-action", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify({ username, password })
    });

    if (response.ok) {
        alert("Registration successful. Please log in.");
        window.location.href = "/login";
    } else {
        alert("Registration failed. Username may already be taken.");
    }
});

async function checkProtectedPage() {
    const token = localStorage.getItem("token");
    if (!token) {
        window.location.href = "/login";
        return;
    }

    const response = await fetch("/protected", {
        method: "GET",
        headers: {
            "Authorization": token
        }
    });

    if (response.redirected) {
        window.location.href = response.url;
    }
}

if (window.location.pathname === "/protected") {
    checkProtectedPage();
}
