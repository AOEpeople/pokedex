import { Form, useSubmit, useTransition } from "@remix-run/react";

interface CardProps {
  id: string;
  name: string;
  isCatched: boolean;
}

export const Card = ({ id, name, isCatched }: CardProps) => {
  const submit = useSubmit();
  const transition = useTransition();
  const imageUrl = `https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/${id}.png`;

  return (
    <Form
      method="post"
      replace={false}
      className={`p-4 bg-gray-100 transition-all duration-300 rounded-lg relative text-center ${
        transition.submission?.formData.get("id") === id ? "opacity-25" : ""
      }`}
    >
      <input type="hidden" name="id" value={id} />
      <input
        id={`${name}-isCatched`}
        className="absolute top-5 right-5 rounded-full border-gray-300 text-red-500 shadow-sm focus:border-red-300 focus:ring focus:ring-offset-0 focus:ring-red-200 focus:ring-opacity-50"
        type="checkbox"
        name="isCatched"
        defaultChecked={isCatched}
        onChange={(e) => submit(e.currentTarget.form)}
      />
      <noscript>
        <button type="submit">Submit</button>
      </noscript>

      <label className="cursor-pointer" htmlFor={`${name}-isCatched`}>
        <img
          width={96}
          height={96}
          className="mx-auto"
          src={imageUrl}
          alt={name}
        />
        <strong>{name}</strong>{" "}
        <span className="text-gray-500 text-sm">#{id}</span>
      </label>
    </Form>
  );
};
