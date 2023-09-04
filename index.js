async function run(){
    let response = await fetch('https://api.freenom.com/v2/service/ping', {
        method: 'GET',
        // mode: 'no-cors',
        headers: {
            "Access-Control-Allow-Origin": "*",
            'content-type': 'application.json'
        }
    })
    console.log(response)
}