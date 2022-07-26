function addItem() {

    let data = document.getElementById("text-input").value
    let url = "http://localhost:8082/tb01"

    fetch(url, {
        method: "POST",
        body: JSON.stringify({text: data}),
        headers: {"Content-type": "application/json; charset=UTF-8"}
    }).then(() => loadItems())

}

function loadItems() {

    let raw = `<table>

                <thead>

                <th id="id">ID</th>
                <th id="text">Text</th>
                <th id="timestamp">Timestamp</th>

                </thead>

            </table>`

    let table = document.getElementsByTagName("table")[0]
    table.innerHTML = raw

    fetch("http://localhost:8081/tb01")
        .then(res => res.json())
        .then(res => res.forEach((row, i) => {

            if ((i % 2) === 0) {

                table.innerHTML = table.innerHTML +
                    `<tr class="even">
                        <td>${row.id}</td>
                        <td>${row.text}</td>
                        <td>${row.date}</td>
                    </tr>`



            } else {

                table.innerHTML = table.innerHTML +
                    `<tr class="odd">
                        <td>${row.id}</td>
                        <td>${row.text}</td>
                        <td>${row.date}</td>
                    </tr>`


            }
        }))

}

function openModal() {
    document.getElementsByClassName("modal")[0].style.display = "block";
}

function hideModal() {
    document.getElementsByClassName("modal")[0].style.display = "none";
}

window.onclick = function(event) {

    let modal = document.getElementsByClassName("modal")[0];

    if (event.target == modal) {
        modal.style.display = "none";
    }
}