class Fakestack < Formula
  desc "High-performance database generator with realistic fake data"
  homepage "https://github.com/0xdps/fake-stack"
  version "2.0.0"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/0xdps/fake-stack/releases/download/v#{version}/fakestack-darwin-arm64"
      sha256 "7746a42485235308c669ca468b6224686dd513fee394e72530e8f7d999357a50"
    else
      url "https://github.com/0xdps/fake-stack/releases/download/v#{version}/fakestack-darwin-amd64"
      sha256 "2daa3716b73a7f8e9393cb780d819178987037d0149da59b0ac45db9d0b57fc4"
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/0xdps/fake-stack/releases/download/v#{version}/fakestack-linux-arm64"
      sha256 "52e5576a25f1d9de496383693e889e4ad866dd3c5ffc8b80c4ec41bc0792b3f9"
    else
      url "https://github.com/0xdps/fake-stack/releases/download/v#{version}/fakestack-linux-amd64"
      sha256 "d046eece8eb50fc799c1ad59788ca469c49e0620f87e66224b75b81a510a9077"
    end
  end

  def install
    bin.install Dir["fakestack-*"].first => "fakestack"
  end

  test do
    system "#{bin}/fakestack", "-d", "."
    assert_predicate testpath/"schema.json", :exist?
  end
end
