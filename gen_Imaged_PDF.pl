use strict;
use warnings;
use PDF::API2::Lite;
use File::Find::Rule;
use File::Basename qw/fileparse/;

print "Dir: ";
chomp( my $dir = <STDIN> );

print "Resolution: ";
chomp( my $resolution = <STDIN> );

# Convert each page of the PDF to PNG
print "\n" . 'Processing...' ."\n";
my $cmd = 'pdf2images_parallel.exe' . ' ' . $dir . ' ' . $resolution;
my $res = `$cmd`;

# Create PDF
my @pdfs = File::Find::Rule->file->name( '*.pdf' )->in($dir);
for  ( @pdfs ){
	my ($basename, $dirname, $ext) = fileparse($_, qr/\..*$/);
	$dirname = $dirname . $basename;

	opendir my $dh, $dirname
	    or die "Can't opendir '$dirname': $!";
	
	my @png = grep -f "$dirname/$_" && /\.png/i, readdir $dh;
	
	closedir $dh
	    or die "Can't closedir '$dirname': $!";
	
	my $pdf = PDF::API2::Lite->new;
	
	for my $png ( sort @png ) {
	
	    my $image = $pdf->image_png( "$dirname/$png" );
	    $pdf->page( $image->width, $image->height );
	    $pdf->image( $image, 0, 0, 1 );
	}
	
	my $result_file = $dirname . '\\' . $basename . $ext;
	$result_file =~ s|/|\\|g;
	$pdf->saveas( $result_file );
	
	print 'Output PDF: ' . $result_file ."\n";
}

print "\nDone!\n";
system('pause > nul');
