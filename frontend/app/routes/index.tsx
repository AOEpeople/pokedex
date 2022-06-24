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

export const loader: LoaderFunction = async () => {
  try {
    const data = await graphqlClient.getPokemon();

    return json(data);
  } catch {
    throw new Response("Es ist ein Fehler aufgetreten", { status: 500 });
  }
};

export const action: ActionFunction = async ({ request }) => {
  const { id, catched } = schema.parse(await request.formData());

  await graphqlClient.setCatched({
    id,
    catched,
  });

  return json({});
};

export default function Index() {
  const { pokemon, total, totalCatched } = useLoaderData<GetPokemonQuery>();

  return (
    <>
      <p className="info">
        {totalCatched} von {total} gefangen!
      </p>

      <div className="list">
        {pokemon.map((p) => (
          <Card key={p.id} {...p} />
        ))}
      </div>
    </>
  );
}
