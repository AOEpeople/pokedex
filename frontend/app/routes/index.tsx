import { ActionFunction, LoaderFunction, json } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";
import { zfd } from "zod-form-data";

import { Card } from "~/components/Card";
import { graphqlClient } from "~/graphql/index.server";
import { GetPokemonQuery } from "~/graphql/sdk";

const schema = zfd.formData({
  id: zfd.numeric(),
  catched: zfd.checkbox(),
});

export const loader: LoaderFunction = async ({}) => {
  return graphqlClient.getPokemon();
};

export const action: ActionFunction = async ({ request }) => {
  const formData = schema.parse(await request.formData());

  // @ todo - replace ith graphql mutation
  console.log(formData);
  return json({});
};

export default function Index() {
  const { pokemon, total, totalCatched } = useLoaderData<GetPokemonQuery>();

  return (
    <div className="container">
      <h1 className="headline">
        AOE <span className="highlight">Pokedex</span>
      </h1>

      <p className="info">
        {totalCatched} von {total} gefangen!
      </p>

      <div className="list">
        {pokemon?.map((p) => (
          <Card
            key={p.id}
            name={p.name}
            catched={p.catched}
            id={p.id.toString()}
          />
        ))}
      </div>
    </div>
  );
}
