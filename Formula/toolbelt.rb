class Toolbelt < Formula
  desc "Toolbelt consolidates scripts into a CLI tool"
  homepage "https://github.com/stianfro/toolbelt"
  url "https://github.com/stianfro/toolbelt/archive/refs/tags/v0.0.0.tar.gz"
  sha256 "0000000000000000000000000000000000000000000000000000000000000000"
  license "MIT"

  depends_on "go" => :build

  def install
    system "go", "build", *std_go_args
  end

  test do
    assert_match "toolbelt", shell_output("#{bin}/toolbelt --help")
  end
end
