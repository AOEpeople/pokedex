import { ActionFunction, LoaderFunction, json } from "@remix-run/node";
import {
  Form,
  useLoaderData,
  useSubmit,
  useTransition,
} from "@remix-run/react";
import { zfd } from "zod-form-data";

interface LoaderData {
  pokemon: Array<{ id: string; name: string; isCatched: boolean }>;
  total: number;
  catched: number;
}

const schema = zfd.formData({
  id: zfd.text(),
  isCatched: zfd.checkbox(),
});

export const loader: LoaderFunction = async ({}) => {
  // @todo - replace with graphql query
  return json<LoaderData>({
    pokemon: [
      { id: "0", name: "Bisasam", isCatched: false },
      { id: "1", name: "Bisaknosp", isCatched: false },
      { id: "2", name: "Bisaflor", isCatched: true },
    ],
    total: 3,
    catched: 1,
  });
};

export const action: ActionFunction = async ({ request }) => {
  const formData = schema.parse(await request.formData());

  // @ todo - replace ith graphql mutation
  console.log(formData);
  return json({});
};

export default function Index() {
  const { pokemon, total, catched } = useLoaderData<LoaderData>();
  const submit = useSubmit();
  const transition = useTransition();

  return (
    <div className="max-w-4xl my-14 mx-auto">
      <h1 className="text-center mb-10 text-2xl">AOE Pokedex</h1>

      <p className="text-right mb-4 text-sm">
        {catched} von {total} gefangen!
      </p>
      <div className="grid gap-4 sm:grid-cols-2 md:grid-cols-3">
        {pokemon.map((p) => (
          <Form
            method="post"
            replace={false}
            className={`p-4 bg-red-100 transition-all duration-300 ${
              transition.submission?.formData.get("id") === p.id
                ? "opacity-25"
                : ""
            }`}
            key={p.id}
          >
            <input type="hidden" name="id" value={p.id} />
            <input
              id={`${p.name}-isCatched`}
              className="mr-2"
              type="checkbox"
              name="isCatched"
              defaultChecked={p.isCatched}
              onChange={(e) => submit(e.currentTarget.form)}
            />
            <noscript>
              <button type="submit">Submit</button>
            </noscript>
            <label className="cursor-pointer" htmlFor={`${p.name}-isCatched`}>
              {p.name} #{p.id}
            </label>
          </Form>
        ))}
      </div>
    </div>
  );
}
