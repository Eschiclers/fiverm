<div id="top"></div>

<!-- PROJECT LOGO -->
<br />

<div align="center">

<img src="/doc/logo.svg">

</div>

  <p align="center">
    FiveM resource manager.
    <br />
    <a href="https://github.com/Eschiclers/fiverm"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/Eschiclers/fiverm">View Demo</a>
    ·
    <a href="https://github.com/Eschiclers/fiverm/issues">Report Bug</a>
    ·
    <a href="https://github.com/Eschiclers/fiverm/issues">Request Feature</a>
  </p>
</div>

<!-- BADGES -->
<div align="center">

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![GPL 3.0][license-shield]][license-url]

</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li>
      <a href="#usage">Usage</a>
      <ul>
        <li><a href="#creating-a-new-project">Creating a new project</a></li>
        <li><a href="#adding-resources-to-the-project">Adding resources to the project</a></li>
      </ul>
    </li>
    <li>
      <a href="#roadmap">Roadmap</a>
      <ul>
        <li><a href="#for-server">For server</a></li>
        <li><a href="#for-resource-creation">For resource creation</a></li>
      </ul>
    </li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#support-me">Support me</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->

## About The Project

fiverm is a FiveM resource manager for your server. It allows you to manage your FiveM resources in a simple and easy way. It is written in [Go](https://github.com/golang/go) for performance and size, and built with [Cobra](https://github.com/spf13/cobra).

<p align="right">(<a href="#top">back to top</a>)</p>


<!-- GETTING STARTED -->

## Getting Started

This is an example of how you may give instructions on setting up your project locally.
To get a local copy up and running follow these simple example steps.

### Installation

For now the only way to install fiverm is from the [actions page](https://github.com/Eschiclers/fiverm/actions) or by downloading it from the [release list](https://github.com/Eschiclers/fiverm/releases).

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- USAGE EXAMPLES -->

## Usage

### Creating a new project

To create a project (resources.json file) you should run the next command in your server folder (along with the server.cfg file):

```console
user@host:~$ fiverm init
```

If you want to create the project again where one already exists, you can use the `--force` flag to overwrite the project. This will delete the project you have created in that folder, but WILL NOT DELETE ANY INSTALLED RESOURCES.

```console
user@host:~$ fiverm init --force
```

### Adding resources to the project

You need to run the commands in the folder where the project is created with the `fiverm init` command

To install resources in a created project, use the install command followed by the username and the repository name as in the example:

```console
user@host:~$ fiverm install eschiclers/zrp_demo
```

If you want the resource to be installed in a custom folder you can use the `--folder` flag and the folder name without `[ ]` as in the example:

```console
user@host:~$ fiverm install eschiclers/zrp_demo --folder zrp
```
And this will install the resource in the `[zrp]` folder.

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- ROADMAP -->

## Roadmap

### For server

- [X] Project creation
- [X] Install resources
  - [X] In custom folder
  - [ ] Install from folder without download
- [ ] Delete resources
- [ ] Update resources

### For resource creation

- [ ] Resource creation

I also want to create a web page as a directory that collects the public resources and thus be able to list them all and create pages such as "the most installed resources".

But for that you need to have servers. Which means spending money. That's why there is a section to support the project through ko-fi.

See the [open issues](https://github.com/Eschiclers/fiverm/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- CONTRIBUTING -->

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#top">back to top</a>)</p>

## Support me
You can support me and this project through ko-fi

[![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/chicle)

<!-- LICENSE -->

## License

Distributed under the [GPL 3.0](https://www.gnu.org/licenses/gpl-3.0.html) license. See `LICENSE` for more information.

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- CONTACT -->

## Contact

Chicle - [@Eschiclers](https://twitter.com/Eschiclers) - [hola@chicle.dev](mailto:hola@chicle.dev) - Telegram: [@Eschiclers](https://t.me/Eschiclers)

Project Link: [https://github.com/Eschiclers/fiverm](https://github.com/Eschiclers/fiverm)

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- ACKNOWLEDGMENTS -->

## Acknowledgments

- [@othneildrew](https://github.com/othneildrew/) for the [README template](https://github.com/othneildrew/Best-README-Template/blob/master/BLANK_README.md)

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[contributors-shield]: https://img.shields.io/github/contributors/Eschiclers/fiverm.svg?style=for-the-badge
[contributors-url]: https://github.com/Eschiclers/fiverm/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/Eschiclers/fiverm.svg?style=for-the-badge
[forks-url]: https://github.com/Eschiclers/fiverm/network/members
[stars-shield]: https://img.shields.io/github/stars/Eschiclers/fiverm.svg?style=for-the-badge
[stars-url]: https://github.com/Eschiclers/fiverm/stargazers
[issues-shield]: https://img.shields.io/github/issues/Eschiclers/fiverm.svg?style=for-the-badge
[issues-url]: https://github.com/Eschiclers/fiverm/issues
[license-shield]: https://img.shields.io/github/license/Eschiclers/fiverm.svg?style=for-the-badge
[license-url]: https://github.com/Eschiclers/fiverm/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/linkedin_username
[product-screenshot]: images/screenshot.png
