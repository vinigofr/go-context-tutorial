# Como contextos funcionam em Go

- Fonte: [tutorialedge.net](https://tutorialedge.net/golang/go-context-tutorial/)
- Traduzido e adaptado por [Vinicius Gouveia](https://linkedin.com/in/vinigofr)

- Última atualização: 27/05/2022 - 00:09


 # O que são contextos?
- De grosso modo, são nossos amigos. Através deles, podemos evitar gastos computacionais desnecessários.
 - Podemos imaginar contextos como uma parcela de informação que será enviada entre diferentes camadas da sua aplicação
 - Essa parcela de informação é criada na borda da aplicação, normalmente quando se recebe uma requisição API
 - Essa parcela de informação é entregue para a camada de serviço ou para a camada de armazenamento

 Duas peças fundamentais para o funcionamento
 1. A habilidade de armazenar informação adicional
 2. A habilidade de controlar cancelamentos


# Bad Practices - Using Contexts For Everything
It should be noted that whilst you can certainly use contexts to pass information between the layers of your application, you absolutely need to use this only for things that truly need to be propagated through.

You shouldn’t use contexts as a bucket for all of your information. It’s a supplementary object that you can store things like request IDs for example or trace IDs which can then be used for logging and tracing purposes.