# fpc-ex
Escreva um programa que implemente o problema clássico de produtor-consumidor. Nesta implementação, você deve receber 3 parâmetros inteiros: i) o número de produtores; ii) o número de consumidores; e iii) o tamanho do buffer. Os produtores produzem duplas de valores inteiros. O primeiro valor é o identificador único daquele produtor e o segundo é o valor de fato produzido (esse segundo elemento produzido deve respeitar uma ordem incremental começando de zero). Cada produtor deverá produzir um número 10 vezes maior do que o tamanho do buffer. Sua implementação deve evitar que produtores trabalham quanto o buffer estiver cheio. De modo análogo, o consumidor não deve trabalhar caso não exista valores no buffer. Um elemento enviado para o buffer não pode ser consumido mais de uma vez. Seu programa deve escrever na saída padrão a seguinte informação, separado por um espaço em branco, quando um valor for consumido: i) o ID único do consumidor; e ii) o valor consumido.