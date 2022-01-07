import React from 'react';
import Layout from '../components/Layout';

export const About = (): JSX.Element => {
  return (
    <Layout
      customMeta={{
        title: 'About - Kaushal Rohit',
      }}
    >
      <h1>About Me</h1>
      <p>
        I&apos;m Kaushal Rohit, a Software Developer, and this is my Personal
        Website. I&apos;m currently working as a Platform Engineer at Benzinga,
        where I build high-throughput and low-latency financial services in
        Golang.
      </p>
      <p>
        When I&apos;m not working, I usually spend my time building things,
        reading books, playing games or watching movies.
      </p>
      Some of the books that I have read and highly recommend are:
      <ul className="list-disc pl-6 my-2">
        <li>Structure and Interpretation of Computer Programs</li>
        <li className="mt-2">Designing Data Intensive Applications</li>
        <li className="mt-2">Effective Python</li>
      </ul>
      Books I&apos;m currently reading:
      <ul className="list-disc pl-6 my-2">
        <li>Algorithm Design Manual</li>
      </ul>
      You can reach out to me at{' '}
      <a href="mailto:rohit.kg98@gmail.com">rohit.kg98@gmail.com</a>
    </Layout>
  );
};

export default About;
