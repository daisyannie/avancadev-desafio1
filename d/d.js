const express = require('express')
const app = express()
const port = 3000

app.get('/coupons/:id', (req, res) => {
    if (req.params.id === 'abc'){
        retorno = 'OK'
    } else{
        retorno = 'NOK'
    }
    console.log(retorno);
    res.send(retorno);    
})

app.listen(port, () => {
  console.log('WS no ar!')
})