curl -s https://learn.01founders.co/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"jsmith\"}}){id}}"}' | cut -d ":" -f4 | cut -b "1-4"





