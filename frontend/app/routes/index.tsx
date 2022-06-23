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
    <div className="max-w-4xl my-14 px-8 mx-auto">
      <h1 className="text-center mb-10 text-5xl font-extrabold tracking-tighter">
        AOE{" "}
        <span className="relative after:absolute after:left-0 after:bottom-[4px] after:-z-10 after:w-full after:h-3 after:bg-red-400">
          Pokedex
        </span>
      </h1>

      <p className="text-right mb-4 text-sm">
        {totalCatched} von {total} gefangen!
      </p>

      <div className="grid gap-4 sm:grid-cols-2 md:grid-cols-3">
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
