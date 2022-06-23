import { Form, useSubmit } from "@remix-run/react";

import { Pokemon } from "~/graphql/sdk";

export const Card = ({ id, name, catched }: Pokemon) => {
  const submit = useSubmit();
  const imageUrl = `https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/${id}.png`;

  return (
    <Form method="post" replace={false} className="card">
      <input type="hidden" name="id" value={id} />
      <input
        id={`${name}-catched`}
        className="checkbox"
        type="checkbox"
        name="catched"
        defaultChecked={catched}
        onChange={(e) => submit(e.currentTarget.form)}
      />
      <noscript>
        <button type="submit">Submit</button>
      </noscript>

      <label className="cursor-pointer" htmlFor={`${name}-catched`}>
        <img
          width={96}
          height={96}
          className="mx-auto"
          src={imageUrl}
          alt={name}
        />
        <strong className="capitalize">{name}</strong>{" "}
        <span className="text-gray-500 text-sm">#{id}</span>
      </label>
    </Form>
  );
};
