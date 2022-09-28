<template>
  <div>
    HELLO

    <!-- <p v-if="$fetchState.pending">Loading....</p>
    <p v-else-if="$fetchState.error">Error while fetching mountains</p>
    <ul v-else>
      <li v-for="(data, index) in data" :key="index">
        {{ data }}
      </li>
    </ul> -->
  </div>
</template>
<script>
  // export default {
  //   data() {
  //     return {
  //       data: []
  //     }
  //   },
    // async fetch() {
    //   this.data = await fetch(
    //     `https://localhost:8080/echo`, 
    //     {
    //       method: 'POST',
    //       mode: 'cors',
    //       headers: {
    //           'Content-Type': 'application/json'
    //       },
    //     }
        
    //   ).then(res => res.json())
    //   console.log(data)
    // }
  // }

const axios = require('axios');

url = "http://localhost:8080/echo"

const config = {
    params: {
    id: 123,
    msg: "TESTTET",
  }
};

axios.get(url, config)
  .then(function (response) {
    // handle success
    console.log(response.data);
  })
  .catch(function (error) {
    // handle error
    console.log(error);
  })
  .then(function () {
    // always executed
  });

// 'https://catfact.ninja/fact'

</script>

