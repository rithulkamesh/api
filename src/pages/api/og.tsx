/* eslint-disable @next/next/no-img-element */
import { ImageResponse } from "@vercel/og";
import type { NextApiRequest, NextApiResponse } from "next";

export const config = {
  runtime: "experimental-edge",
};

const handler = async (req: NextApiRequest, _res: NextApiResponse) => {
  const Inter = await fetch(
    new URL("../../assets/inter.woff", import.meta.url)
  ).then(async (res) => res.arrayBuffer());

  const Hack = await fetch(
    new URL("../../assets/hack.woff", import.meta.url)
  ).then((res) => res.arrayBuffer());

  const { searchParams } = new URL(req.url!);
  const title = searchParams.get("title");
  const top = searchParams.get("top");
  const lg = {
    fontSize: "72px",
    lineHeight: "80px",
    fontWeight: 800,
    fontFamily: "Inter",
    color: "#268bd2",
  };

  const md = {
    fontSize: "62px",
    lineHeight: "70px",
    fontWeight: 900,
    fontFamily: "Inter",
    color: "#268bd2",
  };

  return new ImageResponse(
    (
      <div
        style={{
          height: "100%",
          width: "100%",
          display: "flex",
          justifyContent: "center",
          alignItems: "center",
          color: "white",
          background: "#002b36",
        }}
      >
        <div
          style={{
            display: "flex",
            flexDirection: "column",
            alignItems: "flex-start",
            justifyContent: "space-between",
            width: "1200px",
            height: "630px",
            padding: "80px",
          }}
        >
          <p
            style={{
              fontFamily: "Hack",
              fontSize: "28px",
              marginBottom: "25px",
              color: "#eee8d5",
            }}
          >
            {top}
          </p>

          <h1 style={title!.length < 60 ? lg : md}>{title}</h1>

          <div
            style={{
              display: "flex",
              alignItems: "center",
              justifyContent: "space-between",
              width: "100%",
            }}
          >
            <p
              style={{
                fontFamily: "Hack",
                fontSize: "28px",
                color: "#eee8d5",
              }}
            >
              rithul.dev
            </p>
            <img
              src="https://api.rithul.dev/pfp.png"
              alt="Rithul's avatar"
              width={70}
              height={70}
              style={{
                borderRadius: "100px",
              }}
            />
          </div>
        </div>
      </div>
    ),
    {
      fonts: [
        {
          name: "Inter",
          data: Inter,
          weight: 700,
          style: "normal",
        },
        {
          name: "Hack",
          data: Hack,
          weight: 500,
          style: "normal",
        },
      ],
    }
  );
};

export default handler;
