import Image from "next/image";
import styles from "./page.module.css";
import Link from "next/link";

export default function Home() {
  return (
    <div style={{ display: "flex", flexDirection: "column" }}>
      <Link href="/">Home</Link>
      <Link href="/about">About OK</Link>
      <Link href="/career">Career</Link>
    </div>
  );
}
