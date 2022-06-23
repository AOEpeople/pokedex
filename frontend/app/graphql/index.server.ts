import { GraphQLClient } from "graphql-request";

import { getSdk } from "~/graphql/sdk";

export const graphqlClient = getSdk(
  new GraphQLClient(`http://localhost:3322/graphql`)
);
