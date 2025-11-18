class Fakestack < Formula
  desc "High-performance database generator with realistic fake data"
  homepage "https://github.com/0xdps/fake-stack"
  version "1.1.0"
  license "MIT"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/0xdps/fake-stack/releases/download/v#{version}/fakestack-darwin-arm64"
      sha256 "" # darwin-arm64
    else
      url "https://github.com/0xdps/fake-stack/releases/download/v#{version}/fakestack-darwin-amd64"
      sha256 "" # darwin-amd64
    end
  end

  on_linux do
    if Hardware::CPU.arm?
      url "https://github.com/0xdps/fake-stack/releases/download/v#{version}/fakestack-linux-arm64"
      sha256 "" # linux-arm64
    else
      url "https://github.com/0xdps/fake-stack/releases/download/v#{version}/fakestack-linux-amd64"
      sha256 "" # linux-amd64
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
