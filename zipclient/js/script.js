var form = document.getElementById("form");
var textbox = document.getElementById("textbox");
var button = document.getElementById("button");

var city_name_box = document.getElementById("city_name");


form.addEventListener("submit", function(e) {
    e.preventDefault() //always need this

    city_name = textbox.value
    console.log(city_name)

    // url = "http://localhost:4000/zips/" + city_name
    url = "http://localhost/zips/" + city_name

    fetch(url).then(function(response) {
        return response.text();
    })
        .then(function(data) {
            city_name_box.textContent = data
            console.log(data)
        })
        .catch(function(err) {
            console.error(err)
            alert(err.message);
        })

//     name = textbox.value
//     url = "http://localhost:4000/hello?name=" + name
//     url2 = "http://localhost:4000/memory"

//  fetch(url).then(function(response) {
//         return response.text();
//       })
//         .then(function(data){
//             greeting.textContent = data
//             console.log(data)
//         })
//         .catch(function(err) {
//             console.error(err)
//             alert(err.message);

//         })
//     fetch(url2).then(function(response) {
//         return response.json();
//         })
//         .then(function(data){
//             memory_count.textContent = data.Alloc
//             console.log(data)
//         })
//         .catch(function(err) {
//             console.error(err)
//             alert(err.message);

//         })
  

});
console.log("done")

