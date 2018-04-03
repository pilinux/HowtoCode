/*
 * Inheritance:
 * - Create classes which inherit certain aspects from parent classes.
 */

using System;

namespace Inheritance
{
    class Program
    {
        static void Main(string[] args)
        {
            //Step 2: Create 'vps' variable and call 'Providers' method. Run the program to see the output.
            Vps vps = new Vps();
            vps.Providers();

            //Step 4: Create 'loadBalance' variable and call 'Providers' method. Run the program to see the output.
            LoadBalance loadBalance = new LoadBalance();
            loadBalance.Providers();
            
            //Step 6: Create 'loadBalance2' variable and call both 'Providers' & 'Providers2' methods.
            //Run the program to see the output.
            LoadBalance2 loadBalance2 = new LoadBalance2();
            loadBalance2.Providers();
            loadBalance2.Providers2();
            
            //Step 8: Create 'cdn' variable and call 'CdnProviders' method. Run the program to see the output.
            Cdn cdn = new Cdn();
            cdn.CdnProviders();
            
            //Step 10: Create 'dns' variable and call 'CdnProviders' method. Run the program to see the output.
            Dns dns = new Dns();
            dns.CdnProviders();
        }
    }

    //Step 1: Create a class 'Vps'
    public class Vps
    {
        public void Providers()
        {
            Console.WriteLine("Linode is awesome!");
        }
    }

    //Step 3: Create class 'LoadBalance' which inherits from 'Vps' class
    public class LoadBalance : Vps
    {
        
    }

    //Step 5: Create class 'LoadBalance2' which inherits from 'LoadBalance' class
    public class LoadBalance2 : LoadBalance
    {
        public void Providers2()
        {
           // base.Providers();
            Console.WriteLine("DigitalOcean rocks!");
           // base.Providers();
        }
    }

    //Step 7: Create a parent class 'Cdn'
    public class Cdn
    {
        public virtual void CdnProviders()    //Add 'virtual' keyword so that we can later override the method
        {
            Console.WriteLine("I love Cloudflare.");
        }
    }

    //Step 9: Create class 'Dns' which inherits from 'Cdn' class
    public class Dns : Cdn
    {
        public override void CdnProviders()    //Add 'override' keyword to override the method
        {
            Console.WriteLine("Visit https://1.1.1.1 to learn about Cloudflare public DNS!");    //Overridden
            
            base.CdnProviders();    //To see the actual method without being overridden
        }
    }
}
