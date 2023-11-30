# LumenAI 💡 ![example workflow](https://github.com/SamuelMolling/projeto-integrado/actions/workflows/go.yml/badge.svg)

Bem vindo(a) ao projeto LumenAI!

Antes de começar leia a documentação do projeto [aqui](https://github.com/samuelmolling/projeto-integrado/tree/main/docs/trab2_doc_projeto.html).

## Build e Deploy

Para realizar o build e deploy, após abrir o seu PR irá rodar a action para testar seu código e realizar o build, gerando um artefato. Exemplo pode ver [aqui](https://github.com/SamuelMolling/projeto-integrado/actions/runs/7040459590/job/19161414121)

## Branch e Release

Neste projeto, adotamos a estratégia de branching conhecida como GitFlow para gerenciar nosso fluxo de trabalho de desenvolvimento e lançamento. O GitFlow é um modelo escalável e flexível que nos ajuda a manter a organização e a eficiência no gerenciamento de nossas versões.

### Estrutura de Branches
**Branch main**: Esta é a branch principal. Ela contém o código do estado atual do produto, que é considerado estável e pronto para produção.

**Branch develop**: Esta branch serve como uma área de integração para features. Todo o desenvolvimento de novas funcionalidades inicia a partir da develop e, uma vez finalizadas, são integradas de volta a ela.

**Branches feature**: Para novas funcionalidades, criamos branches a partir da develop. O nome dessas branches segue o padrão feature/nome-da-funcionalidade. Após a conclusão, elas são mescladas de volta à develop.

**Branches release**: Quando estamos prontos para lançar uma nova versão, criamos uma branch release a partir da develop. Qualquer ajuste final para o lançamento é feito nesta branch. Após a conclusão, ela é mesclada à main e à develop.

**Branches hotfix**: Em caso de necessidade de correções urgentes na branch main, usamos as branches hotfix, que são criadas a partir da main. Após a correção, o hotfix é mesclado à main e à develop.

### Política de Releases

**Versionamento**: Utilizamos [SemVer](https://semver.org/) para o versionamento das releases. Cada lançamento é marcado com um número de versão no formato MAJOR.MINOR.PATCH.

**Ciclo de Release**: Planejamos lançamentos regulares para garantir que as atualizações sejam entregues de maneira previsível e consistente.

**Notas de Lançamento**: Com cada release, publicamos notas de lançamento detalhando as novas funcionalidades, melhorias e correções incluídas.


## Issues e Bugs

Para abrir uma Issue ou um Bug, [clique aqui](https://github.com/samuelmolling/projeto-integrado/issues/new).

## Contribuição

Ajude a comunidade tornando este projeto ainda mais incrível. Leia como contribuir clicando **[aqui](https://github.com/samuelmolling/projeto-integrado/tree/main/CONTRIBUTING.md)** e a **[licença](https://github.com/samuelmolling/projeto-integrado/tree/main/LICENSE.md)**. Estou convencido de que juntos alcançaremos coisas incríveis!

## Arquitetura básica do sistema

Você pode ver um diagrama da arquitetura do sistema [neste link](https://github.com/samuelmolling/projeto-integrado/tree/main/docs/images/aws_iot.png)
